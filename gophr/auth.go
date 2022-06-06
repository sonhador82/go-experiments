package main

import (
	"net/http"
)

func AuthenticateRequest(w http.ResponseWriter, r *http.Request) {
	authenticated := false
	if !authenticated {
		http.Redirect(w, r, "/register", http.StatusFound)
	}
}

