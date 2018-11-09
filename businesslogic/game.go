package businesslogic

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type IGame interface {
	CreateNewGameEntry()
	GetGameEntry()
	DeleteGameEntry()
}

// Game model for database
type Game struct {
	gorm.Model
	DevID         int       `gorm:"REFERENCES DEVELOPERS(ID)"`
	PubID         int       `gorm:"REFERENCES PUBLISHERS(ID)"`
	Genre         string    `gorm:"TYPE:VARCHAR(32)"`
	SearchCreated time.Time `gorm:"TYPE:TIMESTAMP;NOT NULL;DEFAULT NOW()"`
	GameName      string    `gorm:"VARCHAR(128);NOT NULL;INDEX"`
	Summary       string    `gorm:"TEXT"`
}

// GameHandler handles the gorm model and its database operations
type GameHandler struct {
	Model *Game
}

func (gh GameHandler) CreateNewGameEntry(conn *gorm.DB) (uint, error) {
	comparison := Game{GameName: gh.Model.GameName} // temporary Game struct
	var err error

	// Check if the game is already in the collection
	conn.Where(&Game{GameName: gh.Model.GameName}).Find(&comparison)
	if comparison.Model.ID != 0 {
		err = errors.New(comparison.GameName + " is already in the collection")
		return 0, err
	}

	// If game isn't in collection, insert the game
	err = conn.Create(&gh.Model).Error
	if err != nil {
		return 0, err
	}

	return gh.Model.ID, nil
}

func (gh GameHandler) GetGameEntry(conn *gorm.DB) (*Game, error) {
	// return error if game is not in the collection
	if conn.Where(&Game{GameName: gh.Model.GameName}).Find(&gh.Model).RecordNotFound() != false {
		err := errors.New(gh.Model.GameName + " is not in the collection")
		return gh.Model, err
	}

	return gh.Model, nil
}

func (gh GameHandler) DeleteGameEntry(conn *gorm.DB) (int, error) {
	// return error if game is not in the collection
	if conn.Where(&Game{GameName: gh.Model.GameName}).Find(&gh.Model).RecordNotFound() != false {
		err := errors.New(gh.Model.GameName + " cannot be deleted because it is not in the collection")
		return 0, err
	}
	conn.Unscoped().Where(&Game{GameName: gh.Model.GameName}).Delete(&gh.Model)

	return 1, nil
}
