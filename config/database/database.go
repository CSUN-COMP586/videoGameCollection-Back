package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const postGresConnectionString = "POSTGRES_CONNECTION"

var dbConnectionString = os.Getenv(postGresConnectionString)

func OpenDatabaseConn() {
	db, err := gorm.Open("postgres", dbConnectionString)
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}

	defer db.Close()
}
