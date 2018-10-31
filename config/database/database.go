package database

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// local testing
const dbConnectionString = "user=vglibdev password=abc123vglib dbname=vglib sslmode=disable"

// GormConn connection available to all packages
var GormConn *gorm.DB
var dialect = "postgres"

func openDatabaseConn() {
	var err error

	GormConn, err = gorm.Open(dialect, dbConnectionString)
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	openDatabaseConn()
}
