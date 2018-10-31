package viewmodel

import (
	"github.com/jinzhu/gorm"
)

// History model for database
type History struct {
	gorm.Model
	GameID    int `gorm:"REFERENCES GAMES(ID)"`
	CharID    int `gorm:"REFERENCES CHARACTERS(ID)"`
	PubID     int `gorm:"REFERENCES PUBLISHERS(ID)"`
	DevID     int `gorm:"REFERENCES DEVELOPERS(ID)"`
	PeopleID  int `gorm:"REFERENCES PEOPLES(ID)"`
	AccountID int `gorm:"REFERENCES ACCOUNTS(ID)"`
}
