package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/videogamelibrary/config/database"
	"github.com/videogamelibrary/config/routes"
)

func main() {
	// database.MigrateDependencyTables()
	// database.MigrateTables()
	defer database.GormConn.Close()

	router := routes.NewCollectionRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	log.Fatal(http.ListenAndServe(port, handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}))(router)))
}
