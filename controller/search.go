package controller

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func SearchForGame(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not supported", 405)
		return
	}

	vars := mux.Vars(r)
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

	searchJSON, err := ioutil.ReadAll(res.Body) // convert reader to bytes
	if err != nil {
		log.Fatal(err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8") // write json to response
	w.WriteHeader(http.StatusOK)
	w.Write(searchJSON)
}
