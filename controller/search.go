package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func SearchForGame(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet { // method check
		http.Error(w, "Method not supported", 405)
		return
	}

	vars := mux.Vars(r) // get route variables
	query := vars["query"]

	client := &http.Client{} // create request
	req, err := http.NewRequest(
		"GET",
		"https://api-endpoint.igdb.com/games/?search="+query+"&fields=id,name",
		nil,
	)
	if err != nil {
		log.Fatal(err)
		return
	}
	req.Header.Set("user-key", "1709f0aad21291fe4c65267aa5141f4e")

	res, err := client.Do(req) // send request
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
