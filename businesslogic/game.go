package businesslogic

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/videogamelibrary/models"
)

type IGameHandler interface {
	CreateNewGameEntry()
	GetGameEntry()
	DeleteGameEntry()
}

// GameHandler handles the gorm model and its database operations
type GameHandler struct {
	Model *models.Game
}

func (handler GameHandler) CreateNewGameEntry(conn *gorm.DB) (uint, error) {
	// check if the game is already in the record under the same account
	check := models.Game{
		AccountID: handler.Model.AccountID,
		GameID:    handler.Model.GameID,
	}
	if conn.Where(&check).Find(&handler.Model).RecordNotFound() != true {
		err := errors.New(handler.Model.GameName + " is in the collection")
		return 0, err
	}

	// otherwise create a new entry
	err := conn.Create(&handler.Model).Error
	if err != nil {
		return 0, err
	}

	return handler.Model.ID, nil
}

func (handler GameHandler) GetGameEntry(conn *gorm.DB, accountID uint) (*[]models.Game, error) {
	var listOfGames []models.Game // slice of game to handle return

	// return error if game is not in the collection
	if conn.Where(&models.Game{AccountID: accountID}).Find(&listOfGames).RecordNotFound() != false {
		err := errors.New("User with ID: " + fmt.Sprint(handler.Model.AccountID) + " is not in the collection")
		return nil, err
	}

	return &listOfGames, nil
}

func (handler GameHandler) DeleteGameEntry(conn *gorm.DB) error {
	game := models.Game{
		AccountID: handler.Model.AccountID,
		GameName:  handler.Model.GameName,
	}

	// return error if game is not in the collection
	if conn.Where(&game).Find(&handler.Model).RecordNotFound() != false {
		err := errors.New(handler.Model.GameName + " cannot be deleted because it is not in the collection")
		return err
	}
	conn.Unscoped().Where(&models.Game{GameName: handler.Model.GameName}).Delete(&handler.Model)

	return nil
}
