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
	GameID        int       `gorm:"NOT NULL; INDEX"`
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

func (handler GameHandler) CreateNewGameEntry(conn *gorm.DB) (uint, error) {
	comparison := Game{GameName: handler.Model.GameName} // temporary Game struct
	var err error

	// Check if the game is already in the collection
	conn.Where(&Game{GameName: handler.Model.GameName}).Find(&comparison)
	if comparison.Model.ID != 0 {
		err = errors.New(comparison.GameName + " is already in the collection")
		return 0, err
	}

	// If game isn't in collection, insert the game
	err = conn.Create(&handler.Model).Error
	if err != nil {
		return 0, err
	}

	return handler.Model.ID, nil
}

func (handler GameHandler) GetGameEntry(conn *gorm.DB) (*Game, error) {
	// return error if game is not in the collection
	if conn.Where(&Game{GameName: handler.Model.GameName}).Find(&handler.Model).RecordNotFound() != false {
		err := errors.New(handler.Model.GameName + " is not in the collection")
		return handler.Model, err
	}

	return handler.Model, nil
}

func (handler GameHandler) DeleteGameEntry(conn *gorm.DB) (int, error) {
	// return error if game is not in the collection
	if conn.Where(&Game{GameName: handler.Model.GameName}).Find(&handler.Model).RecordNotFound() != false {
		err := errors.New(handler.Model.GameName + " cannot be deleted because it is not in the collection")
		return 0, err
	}
	conn.Unscoped().Where(&Game{GameName: handler.Model.GameName}).Delete(&handler.Model)

	return 1, nil
}
