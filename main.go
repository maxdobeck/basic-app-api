package main

import (
	"github.com/maxdobeck/gatekeeper/authentications"
	"net/http"
)

func main() {
	http.HandleFunc("/validate", gatekeeper.ValidSession)
	http.HandleFunc("/login", gatekeeper.Login)
	http.HandleFunc("/logout", gatekeeper.Logout)

	http.ListenAndServe(":3030", nil)
}
