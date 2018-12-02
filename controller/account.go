package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
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

	responseJSON, err := json.Marshal(response) // convert message to json
	if err != nil {
		fmt.Println("error in converting to json")
		log.Fatal(err)
		return
	}

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

	// create handler and account struct, then get account
	account := businesslogic.Account{}
	handler := businesslogic.AccountHandler{
		Model: &account,
	}
	loginStatus, err := handler.GetAccount(database.GormConn, &creds)

	if err != nil {
		fmt.Println("login failure")
		log.Fatal(err)
	}

	response := make(map[string]string)

	// if login failed then encode fail response and write to response writer
	if loginStatus != true {
		response["message"] = err.Error()
		responseJSON, err := json.Marshal(response)
		if err != nil {
			fmt.Println("error encoding loginStatus != true")
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write(responseJSON)
	}

	// otherwise create a jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  handler.Model.ID,
		"exp":  time.Now().Add(time.Minute * 15),
		"role": "user",
		"csrf": os.Getenv("CSRF_KEY"),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_SIGN_KEY")))
	if err != nil {
		fmt.Println("signed string error")
		log.Fatal(err)
	}

	// encode the token and write to response writer
	response["token"] = tokenString
	responseJSON, err := json.Marshal(response)
	if err != nil {
		fmt.Println("error encoding loginStatus == true")
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}
