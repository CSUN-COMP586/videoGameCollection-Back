package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Game model for database
type Game struct {
	gorm.Model
	DevID         int       `gorm:"NOT NULL;REFERENCES DEVELOPERS(ID)"`
	PubID         int       `gorm:"NOT NULL;REFERENCES PUBLISHERS(ID)"`
	Genre         string    `gorm:"TYPE:VARCHAR(32);NOT NULL"`
	SearchCreated time.Time `gorm:"TYPE:TIMESTAMP;NOT NULL;DEFAULT NOW()"`
	GameName      string    `gorm:"VARCHAR(128);NOT NULL;INDEX"`
	Summary       string    `gorm:"TEXT;NOT NULL"`
}
