package main

import (
	"net/http"
	"github.com/maxdobeck/gatekeeper/session_management"
)

func main() {
	http.HandleFunc("/validate", sessions.ValidSession)
	http.HandleFunc("/login", sessions.Login)
	http.HandleFunc("/logout", sessions.Logout)

	http.ListenAndServe(":8080", nil)
}