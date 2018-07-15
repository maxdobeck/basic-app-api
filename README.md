[![Build Status](https://travis-ci.org/maxdobeck/basic-app-api.svg?branch=master)](https://travis-ci.org/maxdobeck/basic-app-api)


# basic-app-api
Service to authenticate users, manage user sessions, and process business logic.

## Installing and Running
```
# Starting local server
$ go run main.go

# Run all tests
$ go test ./...

# Compile and create executable
This will place the executable under go/bin
$ go install

# Run the executable
$ basic-app-api
```

## Background
This is essentially the Model and the Controller in the  MVC for cookie-based sessions that are stored in a Postgresql database.  It shouldn't be too hard to change the cookie store to a Redis database if that's what your needs call for.  Both CSRF protection and sessions were implemented from the Gorilla library of Golang packages.  http://www.gorillatoolkit.org/

The Database management is done essentially by hand, no ORM is in use at the moment.  Setting up the database requires following the instructions for Postgresql database creation and management.  If you'd like to use another database you'll need to change quite a bit in the code.  The db/migrations directory will be executable sql files that create your tables or add migrations.  If you execute them in order starting with `setup.sql` all should be well with your database.

## Dependencies
Go dep is used for dependency management and the /vendor files.  This started as a requirment for Heroku deployment but it looks like Golang may eventually use this tool by default.  Read more here: https://github.com/tools/godep

This was built in parallel with https://github.com/maxdobeck/scheduler-frontend.  At the time of this writing it is a dumb Vue.js Single Page App with very little functionality.  Essentially it has a few protected routes and a log-in-log-out system.  It also requires Vuex.  The real value here is the Fetch requests.  You should be able to essentially copy-paste them from this app to any other app.

## Sources
Started with handy helping from:
* [Go Web Examples/Sessions](https://gowebexamples.com/sessions/)
* [Gorilla/Sessions](https://github.com/gorilla/sessions)
