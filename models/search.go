package models

import "github.com/jinzhu/gorm"

// SearchHistory model for database
type SearchHistory struct {
	gorm.Model
	AccountID uint   `gorm:"NOT NULL;REFERENCES ACCOUNTS(ID)"`
	Query     string `gorm:"TYPE:TEXT;NOT NULL"`
}
