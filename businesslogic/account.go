package businesslogic

import (
	"errors"
	"log"
	"regexp"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type IAccount interface {
	CreateNewAccount()
	GetAccount()
	DeleteAccount()
	VerifyEmail()
	ResetPassword()
}

// Account model for database
type Account struct {
	gorm.Model
	UID           string `gorm:"TYPE:TEXT;UNIQUE`
	Username      string `gorm:"TYPE:VARCHAR(16);UNIQUE;NOT NULL;INDEX"`
	FirstName     string `gorm:"TYPE:VARCHAR(64);NOT NULL"`
	LastName      string `gorm:"TYPE:VARCHAR(64);NOT NULL"`
	DateOfBirth   string `gorm:"TYPE:VARCHAR(12);NOT NULL"`
	Email         string `gorm:"TYPE:TEXT;UNIQUE;NOT NULL"`
	EmailVerified bool   `gorm:"TYPE:BOOLEAN;NOT NULL;DEFAULT FALSE"`
	Password      string `gorm:"TYPE:VARCHAR(64);NOT NULL"`
	RefreshToken  string `gorm:"TYPE:TEXT;UNIQUE"`
}

type Login struct {
	UID      string
	Password string
}

type AccountHandler struct {
	Model *Account
}

func (handler AccountHandler) CreateNewAccount(conn *gorm.DB) (string, bool) {
	var message string
	reUser := regexp.MustCompile("accounts_username_key")
	reEmail := regexp.MustCompile("accounts_email_key")

	handler.Model.Password = hashAndSalt([]byte(handler.Model.Password)) // hash password

	// create new account, return error message if connection error occurs
	if err := conn.Create(&handler.Model).Error; err != nil {
		if reUser.MatchString(err.Error()) == true {
			message = "Username already exists."
		}
		if reEmail.MatchString(err.Error()) == true {
			message = "Email already exists."
		}
		if reUser.MatchString(err.Error()) == true && reEmail.MatchString(err.Error()) == true {
			message = "Username and Email already exists."
		}
		return message, false
	}

	// return success message
	message = "Account successfully created.  Please verify e-mail address."

	return message, true
}

// function to hash and salt password for CreateNewAccount() controller
func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	return string(hash)
}

// GetAccount function that checks user login credentials
func (handler AccountHandler) GetAccount(conn *gorm.DB, creds *Login) (bool, error) {
	// Checks if username is in the database and return model, otherwise return error
	if conn.Where(&Account{UID: creds.UID}).Find(&handler.Model).RecordNotFound() != false {
		err := errors.New("Invalid user identification")
		return false, err
	}

	// compares the hash password with the plain password
	response := comparePasswords(handler.Model.Password, []byte(creds.Password))
	if response != true {
		err := errors.New("Invalid password")
		return false, err
	}

	return response, nil
}

// function to compare verify hash for GetAccount()
func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func (handler AccountHandler) VerifyEmail() {

}

func (handler AccountHandler) VerifyUID(conn *gorm.DB, userID string) (bool, error) {
	if conn.Where(&Account{UID: userID}).Find(&handler.Model).RecordNotFound() != false {
		err := errors.New("Invalid user identification")
		return false, err
	}
	return true, nil
}
