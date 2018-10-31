package businesslogic

import "time"

// Game business logic
type Game struct {
	ID            int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DevID         int
	PubID         int
	Genre         string
	SearchCreated time.Time
	GameName      string
	Summary       string
}

func CreateGameEntry() {

}
