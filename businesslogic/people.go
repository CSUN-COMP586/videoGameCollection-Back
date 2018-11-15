package businesslogic

import "github.com/jinzhu/gorm"

// People model for database
type People struct {
	gorm.Model
	PersonName string `gorm:"TYPE:VARCHAR(32);NOT NULL;INDEX"`
	Gender     bool
	VoiceActed []int `gorm:"TYPE:INTEGER[]"`
	Characters []int `gorm:"TYPE:INTEGER[]"`
}
