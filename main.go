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

	// // testing
	// account := businesslogic.Account{
	// 	Username:      "shortaflip",
	// 	FirstName:     "Kevin",
	// 	LastName:      "Enario",
	// 	DateOfBirth:   time.Now(),
	// 	Email:         "kevin@yahoo.com",
	// 	EmailVerified: false,
	// 	Password:      "1234",
	// }

	// handler := businesslogic.AccountHandler{
	// 	Model: &account,
	// }
}
