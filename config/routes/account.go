package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/videogamelibrary/controller"
)

func AccountRouter() http.Handler {
	router := mux.NewRouter()

	router.Path("/create").Methods(http.MethodPost).HandlerFunc(controller.CreateNewAccount)

	return router
}
