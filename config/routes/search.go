package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/videoGameLibrary/videogamelibrary/controller"
)

func SearchRouter() http.Handler {
	router := mux.NewRouter()

	router.Path("/game/{query}").Methods("GET").HandlerFunc(controller.SearchForGame)

	return router
}
