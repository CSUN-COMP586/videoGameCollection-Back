package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/videogamelibrary/businesslogic"
	"github.com/videogamelibrary/config/database"
	"github.com/videogamelibrary/config/middleware"
)

func GetGameEntry(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get game working"))
}

func AddGameEntry(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not supported", 405)
		return
	}

	response := make(map[string]interface{})

	// verify token and handle false verification status
	verifyStatus, accountID, err := middleware.VerifyToken(r, middleware.App, database.GormConn)
	middleware.HandleFalseVerification(verifyStatus, w, err)

	// decode payload to game struct
	game := businesslogic.Game{}
	if err := json.NewDecoder(r.Body).Decode(&game); err != nil {
		log.Fatal(err)
		return
	}

	// create new game entry to database
	game.AccountID = accountID
	gameHandler := businesslogic.GameHandler{Model: &game}
	gameEntryID, err := gameHandler.CreateNewGameEntry(database.GormConn)
	if err != nil {
		fmt.Println("Error inserting game entry to database: ", err.Error())
	}

	// encode new game entry id and success message
	response["message"] = "Successfully added game to database."
	response["game_entry_id"] = gameEntryID
	responseJSON, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error encoding add game entry success message: ", err.Error())
	}

	// write response
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

func DeleteGameEntry(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete game working"))
}
