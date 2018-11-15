package businesslogic

import "github.com/jinzhu/gorm"

// Developer model for database
type Developer struct {
	gorm.Model
	DevName   string `gorm:"TYPE:VARCHAR(32);NOT NULL;INDEX"`
	DevDesc   string `gorm:"TYPE:TEXT"`
	Country   string `gorm:"TYPE:VARCHAR(32)"`
	Website   string `gorm:"TYPE:VARCHAR(64)"`
	Developed []int  `gorm:"TYPE:INTEGER[]"`
}
