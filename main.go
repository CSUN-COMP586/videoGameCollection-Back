package main

import (
	"log"
	"net/http"

	"github.com/videoGameLibrary/videogamelibrary/config/database"
	"github.com/videoGameLibrary/videogamelibrary/config/routes"
)

func main() {
	defer database.GormConn.Close()

	router := routes.NewCollectionRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
