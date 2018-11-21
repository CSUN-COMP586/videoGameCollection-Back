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
    1.  Create a database called 'vglib'.
    2.  Create a user called 'vglibdev'.
    3.  Create a password for user called 'abc123vglib'.
    4.  Grant all priveleges to user.
  2.  Go version go1.10.3    
    1.  Uncomment 'database.MigrateDependencyTables()' and 'database.MigrateTables()'.
    2.  Only run this once and check database for the tables.

FOLDERS
  1.  The businesslogic folder contains all of the logic regarding the functionality of the application.
  2.  The config folder contains configurations for the database and routing.
  3.  The controller folder contains all of the functions for the endpoints.
  4.  The viewmodel folder contains all of models and functionality for communication with the client side.  

This will be updated as I continue to work on the application.
