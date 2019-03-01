package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/videogamelibrary/controller"
)

// GameRouter - game subroutes
func GameRouter() http.Handler {
	router := mux.NewRouter()

	router.Path("/get").Methods(http.MethodGet).HandlerFunc(controller.GetGameEntry)
	router.Path("/add").Methods(http.MethodPost).HandlerFunc(controller.AddGameEntry)
	router.Path("/delete").Methods(http.MethodDelete).HandlerFunc(controller.DeleteGameEntry)

	return router
}
