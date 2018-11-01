package businesslogic

import "github.com/jinzhu/gorm"

// Developer model for database
type Developer struct {
	gorm.Model
	DevName   string `gorm:"TYPE:VARCHAR(32);NOT NULL;INDEX"`
	DevDesc   string `gorm:"NOT NULL"`
	Country   string `gorm:"TYPE:VARCHAR(32);NOT NULL"`
	Website   string `gorm:"TYPE:VARCHAR(64);NOT NULL"`
	Developed []int  `gorm:"TYPE:INTEGER[];NOT NULL"`
}
