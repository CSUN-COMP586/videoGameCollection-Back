# videogamelibrary

Welcome to the video game library CSUN-COMP586 repository.

The purpose of this project is to learn the processes, patterns, and
school of thought into creating an application using MVC architecture.

Requirements:
  1.  SPA
  2.  MVC
  3.  ORM
  4.  DI
  5.  Auth

To launch the application locally:
  1.  PostGreSQL 9.5.14
    1.  Create a database called 'vglib'
    2.  Create a user called 'vglibdev'
    3.  Create a password for user called 'abc123vglib'
    4.  Grant all priveleges to user
  2.  Go version go1.10.3    
    1.  Uncomment 'database.MigrateDependencyTables()' and 'database.MigrateTables()'
    2.  Only run this once and check database for the tables.

FOLDERS
  1.  The businesslogic folder contains all of the logic regarding the functionality of the application.
  2.  The config folder contains configurations for the database.
  3.  The controller folder contains all of the calls for the endpoints.
  4.  The model folder contains all of the models for the database and functions that deal with calculations.  
  5.  The repository folder contains all of the database logic operations.

This will be updated as I continue to work on the application.
