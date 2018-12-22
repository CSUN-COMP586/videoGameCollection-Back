package businesslogic

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/videogamelibrary/models"
)

type ISearchHistoryHandler interface {
	CreateNewEntry()
}

type SearchHistoryHandler struct {
	Model *models.SearchHistory
}

func (handler SearchHistoryHandler) CreateNewEntry(conn *gorm.DB) {
	if err := conn.Create(&handler.Model).Error; err != nil {
		fmt.Println("Error creating a new entry for search history: ", err.Error())
	}
}
