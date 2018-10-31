package viewmodel

import (
	"github.com/jinzhu/gorm"
)

// Publisher model for database
type Publisher struct {
	gorm.Model
	PubName   string `gorm:"TYPE:VARCHAR(32);NOT NULL;INDEX"`
	Country   string `gorm:"TYPE:VARCHAR(32);NOT NULL"`
	Website   string `gorm:"TYPE:VARCHAR(64);NOT NULL"`
	Published []int  `gorm:"TYPE:INTEGER[];NOT NULL"`
}
