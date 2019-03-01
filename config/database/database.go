package database

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

// GormConn - database connection variable for dependency injection
var GormConn *gorm.DB

func openDatabaseConnection(dialect string, dbConnectionString string) *gorm.DB {
	db, err := gorm.Open(dialect, dbConnectionString)
	if err != nil {
		log.Fatal("Error connecting to the database", err)
	}

	return db
}

// Load configuration files, initialize variables, and open the database connection
func init() {
	err := godotenv.Load("vgConfig.env") // load environment variables
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbName := os.Getenv("DBNAME")
	dbUser := os.Getenv("DBUSER")
	dbPass := os.Getenv("DBPASS")
	dbConnectionString := "user=" + dbUser + " password=" + dbPass + " dbname=" + dbName + " sslmode=disable"
	dialect := os.Getenv("DATABASE")

	GormConn = openDatabaseConnection(dialect, dbConnectionString)
}
