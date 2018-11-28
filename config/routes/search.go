package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/videogamelibrary/controller"
)

func SearchRouter() http.Handler {
	router := mux.NewRouter()

	router.Path("/game/{query}").Methods(http.MethodGet).HandlerFunc(controller.SearchForGame)

	return router
}
