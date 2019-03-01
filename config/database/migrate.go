// Run these functions in respective order if no tables are in the database

package database

import "github.com/videogamelibrary/models"

// MigrateTables - These functions should be used in their respective order
// and should only be ran to create or recreate the tables in the database.
func MigrateTables() {
	GormConn.DropTableIfExists(&models.Account{})
	GormConn.AutoMigrate(&models.Account{})
}

// MigrateDependentTables - as above
func MigrateDependentTables() {
	GormConn.DropTableIfExists(
		&models.Game{},
		&models.SearchHistory{},
	)
	GormConn.AutoMigrate(
		&models.Game{},
		&models.SearchHistory{},
	)
}
