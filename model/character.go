package model

import (
	"github.com/jinzhu/gorm"
)

// Character model for database
type Character struct {
	gorm.Model
	GameID   int    `gorm:"NOT NULL;REFERENCES GAMES(ID)"`
	CharName string `gorm:"TYPE:VARCHAR(32);NOT NULL;INDEX"`
	Gender   int
	Alias    []string `gorm:"TYPE:TEXT[]"`
	Species  int
	PeopleID []int `gorm:"TYPE:INTEGER[]"`
}
