package database

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/videoGameLibrary/videogamelibrary/businesslogic"
)

// local testing
const dbConnectionString = "user=vglibdev password=abc123vglib dbname=vglib sslmode=disable"

var GormConn *gorm.DB

func OpenDatabaseConn() {
	var err error
	var dialect = "postgres"

	GormConn, err = gorm.Open(dialect, dbConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection established")
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
		&businesslogic.History{}, &businesslogic.Search{})
	GormConn.AutoMigrate(
		&businesslogic.Game{}, &businesslogic.Character{},
		&businesslogic.History{}, &businesslogic.Search{})
}

func init() {
	OpenDatabaseConn()
}
