package main

import (
	"github.com/videoGameLibrary/videogamelibrary/config/database"
)

func main() {
	// database.MigrateDependencyTables()
	// database.MigrateTables()

	defer database.GormConn.Close()
}
