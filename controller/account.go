package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/videogamelibrary/businesslogic"
	"github.com/videogamelibrary/config/database"
)

func CreateNewAccount(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost { // method check
		http.Error(w, "Method not supported", 405)
		return
	}

	// decode payload into account struct
	account := businesslogic.Account{}
	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		log.Fatal(err)
		return
	}

	handler := businesslogic.AccountHandler{ // create handler for account struct
		Model: &account,
	}

	// create new account and record message
	message := handler.CreateNewAccount(database.GormConn)
	response := make(map[string]string)
	response["message"] = message

	responseJSON, err := json.Marshal(message) // convert message to json
	if err != nil {
		log.Fatal(err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}
