# videogamelibrary

Welcome to the video game library CSUN-COMP586 repository.

The purpose of this project is to learn the processes, patterns, and
school of thought into creating an application using MVC architecture.

Professor:
  I know you don't read Go professor so here are some notes in reading the language.
  1.  Go doesn't have traditional classes, but instead used structs in combination with
      reciever functions.  Usually declared as such "func (h Handler) DoSomething(n string)"
  2.  The languages convention is to be verbose so each function name should act as the comment.
  3.  Go has a lot of double variable declarations as returns from functions.  They return
      a possible error in this way.
  4.  Any structs using a Go interface only has to have the same reciever functions defined
      in order to fullfil the interface contract.

Requirements:
  1.  SPA
  2.  MVC
  3.  ORM
  4.  DI
  5.  Auth

Notes:
  1.  Firebase and the database are used as dependency injections
  2.  Authentication is handled by firebase
  3.  ORM is handled by GORM package
  4.  I believe the application is fairly decoupled, businesslogic and controllers are separate
      with any configuration dependencies injected in.
  5.  Communication with the front end is handled by Go's native net/http package
  6.  Routing is handled by Gorilla/Mux which is a package built on top of net/http package  

To launch the application locally:
  1.  PostGreSQL 9.5.14
    1.  Create a database called 'vglib'.
    2.  Create a user called 'vglibdev'.
    3.  Create a password for user called 'abc123vglib'.
    4.  Grant all priveleges to user.
  2.  Go version go1.10.3    
    1.  Run "go get -d ./..." to get all dependencies
    2.  In main.go uncomment database.MigrateDependencyTables and database.MigrateDependency
    3.  Run "go run main.go" once to build the tables in postgres, then uncomment the commands
    4.  Run "go run main.go" again to start the server.    

 


