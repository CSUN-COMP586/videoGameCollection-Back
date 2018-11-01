package businesslogic

import (
	"time"

	"github.com/jinzhu/gorm"
)

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
