// The purpose of this file is to migrate all of the models from viewmodel package
// using gorm.  This file is to be used manually until I can figure out
// a more appropriate form of use.

package database

import "github.com/videoGameLibrary/videogamelibrary/businesslogic"

// MigrateDependencyTables function to migrate dependency models to the database
func MigrateDependencyTables() {
	GormConn.DropTableIfExists(
		&businesslogic.Account{}, &businesslogic.Developer{},
		&businesslogic.People{}, &businesslogic.Publisher{})
	GormConn.AutoMigrate(
		&businesslogic.Account{}, &businesslogic.Developer{},
		&businesslogic.People{}, &businesslogic.Publisher{})
}

// MigrateTables function to migrate models to database
func MigrateTables() {
	GormConn.DropTableIfExists(
		&businesslogic.Game{}, &businesslogic.Character{},
		&businesslogic.History{}, &businesslogic.Search{})
	GormConn.AutoMigrate(
		&businesslogic.Game{}, &businesslogic.Character{},
		&businesslogic.History{}, &businesslogic.Search{})
}
