package businesslogic

import (
	"time"

	"github.com/jinzhu/gorm"
)

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

// Connection interface is used for polymorphism, allows for testing that is independent of the connection
type Connection interface {
	CreateNewGameEntry()
}

type GormConnection struct {
	GormConn *gorm.DB
}

func (gc GormConnection) CreateNewGameEntry(model Game) {
	gc.GormConn.Debug().Create(&model)
	gc.GormConn.Debug().Save(&model)
}

// func (gc GormConnection) DeleteNewGAmeEntry(model Game) {
// 	gc.GormConn.Debug().Delete(&model)
// }
