package main

import (
	"fmt"
	"time"

	"github.com/videoGameLibrary/videogamelibrary/businesslogic"
	"github.com/videoGameLibrary/videogamelibrary/config/database"
)

func main() {
	defer database.GormConn.Close()

	gameModel := businesslogic.Game{
		DevID:         10,
		PubID:         20,
		Genre:         "Blah",
		SearchCreated: time.Now(),
		GameName:      "Red",
		Summary:       "You fight red",
	}

	handler := businesslogic.GameHandler{Model: &gameModel}
	res, err := handler.CreateNewGameEntry(database.GormConn)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
