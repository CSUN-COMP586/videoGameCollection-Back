package routes

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func NewCollectionRouter() http.Handler {
	router := mux.NewRouter().StrictSlash(true)
	router.Schemes("https")
	router.HandleFunc("/", temp)

	mount(router, "/game/", GameRouter()) // sub-routes
	mount(router, "/search/", SearchRouter())

	return router
}

func temp(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home"))
}

// function to mount the sub-routes
func mount(r *mux.Router, path string, handler http.Handler) {
	r.PathPrefix(path).Handler(
		http.StripPrefix(
			strings.TrimSuffix(path, "/"),
			handler,
		),
	)
}
