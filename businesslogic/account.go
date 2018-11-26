package businesslogic

import (
	"log"
	"regexp"
	"time"

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
	Username      string    `gorm:"TYPE:VARCHAR(16);UNIQUE;NOT NULL;INDEX"`
	FirstName     string    `gorm:"TYPE:VARCHAR(64);NOT NULL"`
	LastName      string    `gorm:"TYPE:VARCHAR(64);NOT NULL"`
	DateOfBirth   time.Time `gorm:"TYPE:DATE;NOT NULL"`
	Email         string    `gorm:"TYPE:TEXT;UNIQUE;NOT NULL"`
	EmailVerified bool      `gorm:"TYPE:BOOLEAN;NOT NULL;DEFAULT FALSE"`
	Password      string    `gorm:"TYPE:TEXT;NOT NULL"`
}

type AccountHandler struct {
	Model *Account
}

func (handler AccountHandler) CreateNewAccount(conn *gorm.DB) string {
	var message string
	reUser := regexp.MustCompile("accounts_username_key")
	reEmail := regexp.MustCompile("accounts_email_key")

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
		return message
	}

	// return success message
	message = "Account successfully created.  Please verify e-mail address."

	return message
}

// function to hash and salt password for CreateNewAccount()
func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}

	return string(hash)
}

func (handler AccountHandler) GetAccount(conn *gorm.DB) {

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
