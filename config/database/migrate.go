// The purpose of this file is to migrate all of the models from viewmodel package
// using gorm.  This file is to be used manually until I can figure out
// a more appropriate form of use.

package database

import "github.com/videoGameLibrary/videogamelibrary/model"

// MigrateDependencyTables function to migrate dependency models to the database
func MigrateDependencyTables() {
	GormConn.AutoMigrate(
		&model.Account{}, &model.Developer{},
		&model.People{}, &model.Publisher{})
}

// MigrateTables function to migrate models to database
func MigrateTables() {
	GormConn.AutoMigrate(
		&model.Game{}, &model.Character{},
		&model.History{}, &model.Search{})
}
