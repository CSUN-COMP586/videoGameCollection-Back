package businesslogic

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// Search model for database
type SearchHistory struct {
	gorm.Model
	AccountID uint   `gorm:"NOT NULL;REFERENCES ACCOUNTS(ID)"`
	Query     string `gorm:"TYPE:TEXT;NOT NULL"`
}

type SearchHistoryHandler struct {
	Model *SearchHistory
}

func (handler SearchHistoryHandler) CreateNewEntry(conn *gorm.DB) {
	if err := conn.Create(&handler.Model).Error; err != nil {
		fmt.Println("Error creating a new entry for search history: ", err.Error())
	}
}
