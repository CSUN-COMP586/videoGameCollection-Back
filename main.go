package main

import (
	"github.com/videoGameLibrary/videogamelibrary/config/database"
)

func main() {
	defer database.GormConn.Close()
}
