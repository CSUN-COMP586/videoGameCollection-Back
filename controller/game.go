package controller

import (
	"net/http"
)

func SearchForGame(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Search working"))
}

func GetGameEntry(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get game working"))
}

func AddGameEntry(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Add game working"))
}

func DeleteGameEntry(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete game working"))
}
