package main

import (
	"net/http"
	"github.com/maxdobeck/gatekeeper/gatekeeper"
)

func main() {
	http.HandleFunc("/validate", gatekeeper.ValidSession)
	http.HandleFunc("/login", gatekeeper.Login)
	http.HandleFunc("/logout", gatekeeper.Logout)

	http.ListenAndServe(":3030", nil)
}