package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/videogamelibrary/businesslogic"
	"github.com/videogamelibrary/config/database"
	"github.com/videogamelibrary/config/middleware"
	"github.com/videogamelibrary/models"
)

func SearchForGame(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet { // method check
		http.Error(w, "Method not supported", 405)
		return
	}

	// verify token, return status of verification, users account id, and potential errors
	verifyStatus, accountID, err := middleware.VerifyToken(r, middleware.App, database.GormConn)

	// if status is false, send response to front end
	middleware.HandleFalseVerification(verifyStatus, w, err)

	// get route variables
	vars := mux.Vars(r)
	query := vars["query"]

	// create search history struct and insert to database
	searchHistory := models.SearchHistory{
		AccountID: accountID,
		Query:     query,
	}
	searchHistoryHandler := businesslogic.SearchHistoryHandler{
		Model: &searchHistory,
	}
	searchHistoryHandler.CreateNewEntry(database.GormConn)

	// create request to IGDB API
	client := &http.Client{}
	req, err := http.NewRequest(
		http.MethodGet,
		"https://api-endpoint.igdb.com/games/?search="+query+"&fields=id,name,summary,cover",
		nil,
	)
	if err != nil {
		log.Fatal(err)
		return
	}
	req.Header.Set("user-key", os.Getenv("IGDB_USER_KEY"))

	// send request to IGDB API
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return
	}

	// decode response body to an array of dict
	search := make([]map[string]interface{}, 1)
	err = json.NewDecoder(res.Body).Decode(&search)
	if err != nil {
		log.Fatal(err)
		return
	}

	// convert to a json
	searchJSON, err := json.Marshal(search)
	if err != nil {
		log.Fatal(err)
		return
	}

	// write json to response
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(searchJSON)
}
