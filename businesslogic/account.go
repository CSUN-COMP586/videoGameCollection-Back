package businesslogic

import (
	"time"

	"github.com/jinzhu/gorm"
)

type IAccount interface {
	CreateNewAccount()
	GetAccount()
	DeleteAccount()
	VerifyEmail()
	ResetPassword()
	HashPassword()
}

// Account model for database
type Account struct {
	gorm.Model
	Username      string     `gorm:"TYPE:VARCHAR(16);NOT NULL;INDEX"`
	FirstName     string     `gorm:"TYPE:VARCHAR(64);NOT NULL"`
	LastName      string     `gorm:"TYPE:VARCHAR(64);NOT NULL"`
	DateOfBirth   *time.Time `gorm:"TYPE:DATE;NOT NULL"`
	Email         string     `gorm:"TYPE:TEXT;UNIQUE;NOT NULL"`
	EmailVerified bool       `gorm:"TYPE:BOOLEAN;NOT NULL;DEFAULT FALSE"`
	HashAlgorithm string     `gorm:"TYPE:TEXT"`
	Password      string     `gorm:"TYPE:VARCHAR(16);NOT NULL"`
	PasswordSalt  string     `gorm:"TYPE:BYTEA"`
	PasswordHash  string     `gorm:"TYPE:BYTEA"`
}

type AccountHandler struct {
	Model *Account
}

func (ah AccountHandler) CreateNewAccount(conn *gorm.DB) string {
	var message string

	// return string if username is taken
	if conn.Where(&Account{Username: ah.Model.Username}).Find(&ah.Model).RecordNotFound() != true {
		message = "Username already taken."

		return message
	}

	// return string if email is taken
	if conn.Where(&Account{Username: ah.Model.Email}).Find(&ah.Model).RecordNotFound() != true {
		message = "E-mail address already taken."

		return message
	}

	return message
}

func createHash() {

}
