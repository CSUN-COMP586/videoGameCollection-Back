package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/videoGameLibrary/videogamelibrary/controller"
)

func GameRouter() http.Handler {
	router := mux.NewRouter()

	router.Path("/get").Methods("GET").HandlerFunc(controller.GetGameEntry)
	router.Path("/add").Methods("POST").HandlerFunc(controller.AddGameEntry)
	router.Path("/delete").Methods("DELETE").HandlerFunc(controller.DeleteGameEntry)

	return router
}
