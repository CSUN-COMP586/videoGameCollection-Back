package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/videogamelibrary/businesslogic"
	"github.com/videogamelibrary/config/database"
	"github.com/videogamelibrary/config/middleware"
	"github.com/videogamelibrary/models"
)

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
	gameModel := models.Game{}
	if err := json.NewDecoder(r.Body).Decode(&gameModel); err != nil {
		fmt.Println("Error decoding body to game struct: ", err.Error())
		return
	}

	// create new game entry to database
	gameModel.AccountID = accountID
	gameHandler := businesslogic.GameHandler{Model: &gameModel}
	gameEntryID, err := gameHandler.CreateNewGameEntry(database.GormConn)

	// if error in inserting game to database occurs
	if err != nil {
		fmt.Println("Error inserting game entry to database: ", err.Error())
		response["message"] = err.Error()
	} else {
		response["message"] = "Successfully added game to database."
		response["game_entry_id"] = gameEntryID
	}
	responseJSON, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error encoding add game entry success message: ", err.Error())
	}

	// write response
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

func GetGameEntry(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not supported", 405)
		return
	}

	response := make(map[string]interface{})

	// verify token and handle false verification status
	verifyStatus, accountID, err := middleware.VerifyToken(r, middleware.App, database.GormConn)
	middleware.HandleFalseVerification(verifyStatus, w, err)

	gameModel := models.Game{}
	gameHandler := businesslogic.GameHandler{Model: &gameModel}
	listOfGames, err := gameHandler.GetGameEntry(database.GormConn, accountID)
	if err != nil {
		fmt.Println("Error getting games from the database: ", err.Error())
		response["response"] = err.Error()
	} else {
		response["response"] = "Success"
		response["data"] = listOfGames
	}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error encoding JSON games response: ", err.Error())
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

func DeleteGameEntry(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete game working"))
}
