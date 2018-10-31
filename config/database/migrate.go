// The purpose of this file is to migrate all of the models from viewmodel package
// using gorm.  This file is to be used manually until I can figure out
// a more appropriate form of use.

package database

import (
	"github.com/videoGameLibrary/videogamelibrary/viewmodel"
)

// MigrateDependencyTables function to migrate dependency models to the database
func MigrateDependencyTables() {
	GormConn.AutoMigrate(
		&viewmodel.Account{}, &viewmodel.Developer{},
		&viewmodel.People{}, &viewmodel.Publisher{})
}

// MigrateTables function to migrate models to database
func MigrateTables() {
	GormConn.AutoMigrate(
		&viewmodel.Game{}, &viewmodel.Character{},
		&viewmodel.History{}, &viewmodel.Search{})
}
