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

func openDatabaseConn() {
	var err error

	GormConn, err = gorm.Open("postgres", dbConnectionString)
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	openDatabaseConn()
}
