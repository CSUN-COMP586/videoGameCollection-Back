# videogamelibrary

Welcome to the video game library CSUN-COMP586 repository.

The purpose of this project is to learn the processes, patterns, and
school of thought into creating an application using MVC architecture and SOLID principles.

### Update
Since the inception of this project, I have learned many more things about back-end 
development, and as such, will continue to polish this project with the main purpose of
being an example of my back-end work.

### Development 
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
