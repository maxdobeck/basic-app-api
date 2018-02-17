# Gatekeeper
Service to authenticate users and validate sessions.

## Installing and Running
```
# Starting local server
go run main.go

# Useful URLs
/validate - Checks if the supplied session is active
/login - Creates a session and returns an HTTP Only cookie
/logout - Destroys the supplied session
```

## Sources
Started with handy helping from:
* [Go Web Examples/Sessions](https://gowebexamples.com/sessions/)
* [Gorilla/Sessions](https://github.com/gorilla/sessions)