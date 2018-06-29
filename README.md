[![Build Status](https://travis-ci.org/maxdobeck/gatekeeper.svg?branch=master)](https://travis-ci.org/maxdobeck/gatekeeper)
# Gatekeeper
Service to authenticate users, create sessions, and validate sessions.

## Installing and Running
```
# Starting local server
$ go run main.go

# Run all tests
$ go test ./tests/*

# Compile and create executable
This will place the executable under go/bin
$ go install

# Run the executable
$ gatekeeper

# Useful URLs
/validate - Checks if the supplied session is active
/login - Creates a session and returns an HTTP Only cookie
/logout - Destroys the supplied session
```

## Sources
Started with handy helping from:
* [Go Web Examples/Sessions](https://gowebexamples.com/sessions/)
* [Gorilla/Sessions](https://github.com/gorilla/sessions)
