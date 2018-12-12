package businesslogic

import (
	"errors"

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
	AccountID uint   `gorm:"NOT NULL; REFERENCES ACCOUNTS(ID)`
	GameID    int    `gorm:"NOT NULL; INDEX"`
	GameName  string `gorm:"VARCHAR(128);NOT NULL;INDEX"`
	Summary   string `gorm:"TEXT"`
	ImageURL  string `gorm:"TYPE:TEXT`
}

// GameHandler handles the gorm model and its database operations
type GameHandler struct {
	Model *Game
}

func (handler GameHandler) CreateNewGameEntry(conn *gorm.DB) (uint, error) {
	err := conn.Create(&handler.Model).Error
	if err != nil {
		return 0, err
	}

	return handler.Model.ID, nil
}

func (handler GameHandler) GetGameEntry(conn *gorm.DB) (*Game, error) {
	// return error if game is not in the collection
	if conn.Where(&Game{GameID: handler.Model.GameID}).Find(&handler.Model).RecordNotFound() != false {
		err := errors.New(handler.Model.GameName + " is not in the collection")
		return nil, err
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
