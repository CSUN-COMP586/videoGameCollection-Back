package models

import "github.com/jinzhu/gorm"

// Game model for database
type Game struct {
	gorm.Model
	AccountID uint   `gorm:"NOT NULL; REFERENCES ACCOUNTS(ID)`
	GameID    int    `gorm:"NOT NULL; INDEX"`
	GameName  string `gorm:"VARCHAR(128);NOT NULL;INDEX"`
	Summary   string `gorm:"TEXT"`
	ImageURL  string `gorm:"TYPE:TEXT`
}
