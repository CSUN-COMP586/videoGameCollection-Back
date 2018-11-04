package businesslogic

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Handler interface {
	CreateNewGameEntry()
}

// Game model for database
type Game struct {
	gorm.Model
	DevID         int       `gorm:"REFERENCES DEVELOPERS(ID)"`
	PubID         int       `gorm:"REFERENCES PUBLISHERS(ID)"`
	Genre         string    `gorm:"TYPE:VARCHAR(32);NOT NULL"`
	SearchCreated time.Time `gorm:"TYPE:TIMESTAMP;NOT NULL;DEFAULT NOW()"`
	GameName      string    `gorm:"VARCHAR(128);NOT NULL;INDEX"`
	Summary       string    `gorm:"TEXT;NOT NULL"`
}

// GameHandler handles the gorm model that is to be manipulated
type GameHandler struct {
	Model *Game
}

func (gh GameHandler) CreateNewGameEntry(conn *gorm.DB) (uint, error) {
	err := conn.Create(&gh.Model).Error
	if err != nil {
		return 0, err
	}
	return gh.Model.ID, nil
}

// func (gc GormConnection) DeleteNewGAmeEntry(model Game) {
// 	gc.GormConn.Debug().Delete(&model)
// }
