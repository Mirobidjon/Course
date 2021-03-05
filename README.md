#Course

This project provides the following features right out of the box
* RESTful endpoints in the widely accepted format
* Standard CRUD operations of a database table
* JWT-based authentication
* Environment dependent application configuration management
* Structured logging with contextual information
* Error handling with proper error response generation
* Database migration
* Data validation

## Getting Started
If this is your first time encountering Go, please follow [the instructions](https://golang.org/doc/install) to
install Go on your computer.

[Docker](https://www.docker.com/get-started) is also needed if you want to try the project without setting up your
own database server.

Run This Project
```shell
# download the starter kit
git clone 

cd go-rest-api

# start a PostgreSQL database server in a Docker container
make createdb

#migrate a database
make migrateup

# run the RESTful API server
make run

#and run 
make run

```

All users can register themselves role

##Roles 
* Students
* Teachers
* Directors

Student can 
* Create book           ("/api/book") POST
* READ a book           ("/api/book/:id") GET
* Read all books        ("api/book") GET
* Update book           ("api/book/:id") PUT
* Delete book           ("api/book/:id) DELETE
* Read himself group course ("api/course/:id") GET
* Read himself group all course ("api/course/:id") GET
* Read all directors    ("/api/director") GET
* Read director         ("/api/director/:id") GET
* Read all teachers     ("/api/teacher") GET
* Read teacher          ("/api/teacher/:id") GET
* Read all students     ("/api/student") GET
* Read student          ("/api/student/:id") GET

Teachers can 
* Create group course   ("/api/course") POST
* Read himself course       ("/api/course/:id") GET
* Read himself all course   ("/api/course/") GET
* Update himself course     ("/api/course/:id") PUT
* Delete course         ("/api/course/:id") DELETE
* Read all directors    ("/api/director") GET
* Read director         ("/api/director/:id") GET
* Read all teachers     ("/api/teacher") GET
* Read teacher          ("/api/teacher/:id") GET
* Read all students     ("/api/student") GET
* Read student          ("/api/student/:id") GET
* Delete student         ("/api/student/:id") DELETE

Directors can 
* Read all course   ("/api/course/") GET
* Read all directors    ("/api/director") GET
* Read director         ("/api/director/:id") GET
* Read all teachers     ("/api/teacher") GET
* Read teacher          ("/api/teacher/:id") GET
* Read all students     ("/api/student") GET
* Read student          ("/api/student/:id") GET
* Update himself        ("/api/director) PUT
* Delete teacher        ("/api/teacher/:id") DELETE