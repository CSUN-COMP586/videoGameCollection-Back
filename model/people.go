package model

import (
	"github.com/jinzhu/gorm"
)

// People model for database
type People struct {
	gorm.Model
	PersonName string `gorm:"TYPE:VARCHAR(32);NOT NULL;INDEX"`
	Gender     bool   `gorm:"NOT NULL"`
	VoiceActed []int  `gorm:"TYPE:INTEGER[];NOT NULL"`
	Characters []int  `gorm:"TYPE:INTEGER[];NOT NULL"`
}
