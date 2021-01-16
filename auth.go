package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	_ = r.ParseForm()

	if r.FormValue("user") == authUser && r.FormValue("password") == authPassword {
		// Delegate request to the given handle
		w.Header().Set("Authorization", "Basic a29zbW9zOkFidWJ1QWVBa2FrMzIy")
		w.Write([]byte(`{"status": "ok", "msg": "you authenticated"}`))
	} else {
		// Request Basic Authentication otherwise
		w.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
	}
}

func BasicAuth(h httprouter.Handle, requiredUser string, requiredPassword string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		// Get the Basic Authentication credentials
		user, password, hasAuth := r.BasicAuth()

		if hasAuth && user == requiredUser && password == requiredPassword {
			// Delegate request to the given handle
			h(w, r, ps)
		} else {
			// Request Basic Authentication otherwise
			w.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}
	}
}