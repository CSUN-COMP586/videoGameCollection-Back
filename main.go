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

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}))(router)))

	// testing
	// creds := businesslogic.Login{
	// 	Username: "breh",
	// 	Password: "blah123",
	// }

	// account := businesslogic.Account{}

	// handler := businesslogic.AccountHandler{
	// 	Model: &account,
	// }

	// loginStatus, err := handler.GetAccount(database.GormConn, &creds)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(loginStatus)

}
