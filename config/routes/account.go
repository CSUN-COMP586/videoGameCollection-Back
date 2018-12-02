package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/videogamelibrary/controller"
)

func AccountRouter() http.Handler {
	router := mux.NewRouter()

	router.Path("/register").Methods(http.MethodPost).HandlerFunc(controller.CreateNewAccount)
	router.Path("/login").Methods(http.MethodPost).HandlerFunc(controller.Login)

	return router
}
