package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/videoGameLibrary/videogamelibrary/model"
)

const (
	nameQuery = "GAME_NAME = ?"
	idQuery   = "ID = ?"
)

type GameRepository struct {
	Database *gorm.DB
}

func (repo GameRepository) AddGameEntry(game model.Game) {
	if repo.Database == nil {
		fmt.Println("No database")
	}
	if repo.Database.NewRecord(game) == true {
		res := repo.Database.Create(&game)
		if res != nil {
			fmt.Println(res)
		}
	}
}

func (repo GameRepository) GetGameEntryByName(name string) {
	if repo.Database == nil {
		fmt.Println("No database")
	}
	res := repo.Database.Where(nameQuery, name)
	if res != nil {
		fmt.Println(res)
	}
}

func (repo GameRepository) GetGameEntryByID(id string) {
	if repo.Database == nil {
		fmt.Println("No database")
	}
	res := repo.Database.Where(idQuery, id)
	if res != nil {
		fmt.Println(res)
	}
}
