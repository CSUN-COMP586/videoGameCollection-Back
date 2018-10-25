package viewmodel

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Account struct {
	gorm.Model
	Username    string     `json:"username"`
	FirstName   string     `json:"firstname"`
	LastName    string     `json:"lastname"`
	DateOfBirth *time.Time `json:"dateofbirth"`
	Email       string     `json:"email"`
	Password    string     `json:"password"`
}
