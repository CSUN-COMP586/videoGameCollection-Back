package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/videogamelibrary/config/database"
	"github.com/videogamelibrary/config/routes"
)

func main() {
	// database.MigrateDependencyTables()
	// database.MigrateTables()
	defer database.GormConn.Close()

	router := routes.NewCollectionRouter()

	// provision production port

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}))(router)))
}
