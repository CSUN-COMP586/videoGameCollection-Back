package database

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/videogamelibrary/businesslogic"
	"github.com/videogamelibrary/config/middleware"
)

// initialize environment variables
var (
	DBNAME             = middleware.GetEnv("DBNAME", "vglib")
	DBUSER             = middleware.GetEnv("DBUSER", "vglibdev")
	DBPASS             = middleware.GetEnv("DBPASS", "abc123vglib")
	dbConnectionString = "user=" + DBUSER + " password=" + DBPASS + " dbname=" + DBNAME + " sslmode=disable"
	dialect            = "postgres"
)

// this pointer variable is initialized with the database connection and can then be called
// for a dependency injection in other packages
var GormConn *gorm.DB

func init() {
	GormConn = openDatabaseConnection(dialect, dbConnectionString)
}

// private functions for connecting to the database
func openDatabaseConnection(dialect string, dbConnectionString string) *gorm.DB {
	var db *gorm.DB
	var err error

	db, err = gorm.Open(dialect, dbConnectionString)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

// MigrateDependencyTables - These functions should be used in their respective order
// and should only be ran to create or recreate the tables in the database.
func MigrateDependencyTables() {
	GormConn.DropTableIfExists(
		&businesslogic.Account{}, &businesslogic.Developer{},
		&businesslogic.People{}, &businesslogic.Publisher{})
	GormConn.AutoMigrate(
		&businesslogic.Account{}, &businesslogic.Developer{},
		&businesslogic.People{}, &businesslogic.Publisher{})
}

// MigrateTables - as above
func MigrateTables() {
	GormConn.DropTableIfExists(
		&businesslogic.Game{}, &businesslogic.Character{},
		&businesslogic.History{}, &businesslogic.SearchHistory{})
	GormConn.AutoMigrate(
		&businesslogic.Game{}, &businesslogic.Character{},
		&businesslogic.History{}, &businesslogic.SearchHistory{})
}
