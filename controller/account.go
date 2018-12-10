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

	accountHandler := businesslogic.AccountHandler{ // create accountHandler for account struct
		Model: &account,
	}

	// create new account and record message
	message, status := accountHandler.CreateNewAccount(database.GormConn)
	response := make(map[string]interface{})
	response["message"] = message
	response["status"] = status

	responseJSON, err := json.Marshal(response) // convert message to json
	if err != nil {
		fmt.Println("error in converting to json")
		log.Fatal(err)
		return
	}

	// email verification

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost { // method check
		http.Error(w, "Method not supported", 405)
		return
	}

	// decode credentials into login struct
	creds := businesslogic.Login{}
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		log.Fatal(err)
		return
	}

	// decode jwt and assign uid to login struct
	idToken := r.Header.Get("authorization")
	authModel := middleware.Auth{Token: idToken}
	authHandler := middleware.AuthHandler{Model: &authModel}
	creds.UID = authHandler.VerifyTokenAndReturnUID()

	// create account handler and account struct, then get account to login
	account := businesslogic.Account{}
	accountHandler := businesslogic.AccountHandler{
		Model: &account,
	}
	loginStatus, err := accountHandler.GetAccount(database.GormConn, &creds)
	if err != nil {
		fmt.Println("Error during server login: " + err.Error())
		log.Fatal(err)
	}

	response := make(map[string]interface{})

	// if login failed then encode fail response and write to response writer
	if loginStatus != true {
		response["message"] = err.Error()
		responseJSON, err := json.Marshal(response)
		if err != nil {
			fmt.Println("Error encoding login failure status: " + err.Error())
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write(responseJSON)
	}

	// encode message and write to response writer
	response["message"] = loginStatus
	responseJSON, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error encoding login success status: " + err.Error())
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}
