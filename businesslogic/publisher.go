package businesslogic

import "github.com/jinzhu/gorm"

// Publisher model for database
type Publisher struct {
	gorm.Model
	PubName   string `gorm:"TYPE:VARCHAR(32);NOT NULL;INDEX"`
	Country   string `gorm:"TYPE:VARCHAR(32)"`
	Website   string `gorm:"TYPE:VARCHAR(64)"`
	Published []int  `gorm:"TYPE:INTEGER[]"`
}
