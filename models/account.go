package models

import "github.com/jinzhu/gorm"

// Account model for database
type Account struct {
	gorm.Model
	UID           string `gorm:"TYPE:TEXT;UNIQUE"`
	Username      string `gorm:"TYPE:VARCHAR(16);UNIQUE;NOT NULL;INDEX"`
	FirstName     string `gorm:"TYPE:VARCHAR(64);NOT NULL"`
	LastName      string `gorm:"TYPE:VARCHAR(64);NOT NULL"`
	DateOfBirth   string `gorm:"TYPE:VARCHAR(12);NOT NULL"`
	Email         string `gorm:"TYPE:TEXT;UNIQUE;NOT NULL"`
	EmailVerified bool   `gorm:"TYPE:BOOLEAN;NOT NULL;DEFAULT FALSE"`
	Password      string `gorm:"TYPE:VARCHAR(64);NOT NULL"`
	RefreshToken  string `gorm:"TYPE:TEXT;UNIQUE"`
}
