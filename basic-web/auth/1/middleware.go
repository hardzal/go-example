package main

import "net/http"

// USERNAME constants
const USERNAME = "batman"

// PASSWORD constants
const PASSWORD = "secret"

// Auth func
func Auth(w http.ResponseWriter, r *http.Request) bool {
	username, password, ok := r.BasicAuth()
	if !ok {
		w.Write([]byte("Something went wrong"))
		return false
	}

	isValid := (username == USERNAME) && (password == PASSWORD)
	if !isValid {
		w.Write([]byte("wrong username/password"))
		return false
	}

	return true
}

// AllowOnlyGET func
func AllowOnlyGET(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != "GET" {
		w.Write([]byte("Only Get is Allowed"))
		return false
	}

	return true
}
