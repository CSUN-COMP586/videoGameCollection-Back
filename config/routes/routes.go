package routes

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// NewCollectionRouter - initializes the router and mounts the sub routes
func NewCollectionRouter() http.Handler {
	router := mux.NewRouter().StrictSlash(true)
	router.Schemes("https")
	router.HandleFunc("/", home)

	mount(router, "/api/game/", GameRouter())
	mount(router, "/api/search/", SearchRouter())
	mount(router, "/api/account/", AccountRouter())

	return router
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home"))
}

func mount(r *mux.Router, path string, handler http.Handler) {
	r.PathPrefix(path).Handler(
		http.StripPrefix(
			strings.TrimSuffix(path, "/"),
			handler,
		),
	)
}
