package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/videogamelibrary/businesslogic"
	"github.com/videogamelibrary/config/database"
	"github.com/videogamelibrary/config/middleware"
)

func SearchForGame(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet { // method check
		http.Error(w, "Method not supported", 405)
		return
	}

	// verify token and return UID
	idToken := r.Header.Get("authorization")
	authModel := middleware.Auth{Token: idToken}
	authHandler := middleware.AuthHandler{Model: &authModel}
	UID := authHandler.VerifyTokenAndReturnUID(middleware.App)
	response := make(map[string]interface{})

	// verify UID and return true or false
	account := businesslogic.Account{}
	accountHandler := businesslogic.AccountHandler{Model: &account}
	verifyStatus, err := accountHandler.VerifyUID(database.GormConn, UID)
	if err != nil {
		fmt.Println("Verify UID error: ", err.Error())
	}

	// if status is false, send response to front end
	if verifyStatus != true {
		response["message"] = "UID not verified." + err.Error()
		responseJSON, err := json.Marshal(response)
		if err != nil {
			fmt.Println("Error encoding search controller UID verification: " + err.Error())
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write(responseJSON)
	}

	vars := mux.Vars(r) // get route variables
	query := vars["query"]

	client := &http.Client{} // create request to IGDB API
	req, err := http.NewRequest(
		http.MethodGet,
		"https://api-endpoint.igdb.com/games/?search="+query+"&fields=id,name",
		nil,
	)
	if err != nil {
		log.Fatal(err)
		return
	}
	req.Header.Set("user-key", os.Getenv("IGDB_USER_KEY"))

	res, err := client.Do(req) // send request to IGDB API
	if err != nil {
		log.Fatal(err)
		return
	}

	search := make([]map[string]interface{}, 1) // decode response body to an array of dict
	err = json.NewDecoder(res.Body).Decode(&search)
	if err != nil {
		log.Fatal(err)
		return
	}

	searchJSON, err := json.Marshal(search) // convert to a json
	if err != nil {
		log.Fatal(err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8") // write json to response
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(searchJSON)
}
