package businesslogic

import (
	"time"
)

type Account struct {
	ID            int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Username      string
	FirstName     string
	LastName      string
	DateOfBirth   time.Time
	Email         string
	EmailVerified string
	HashAlgorithm string
	Password      string
	PasswordSalt  string
	PasswordHash  string
}
