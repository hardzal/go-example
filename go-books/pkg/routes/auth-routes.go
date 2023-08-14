package routes

import (
	"github.com/gorilla/mux"
	"github.com/hardzal/go-example/go-books/pkg/controllers/authcontrollers"
)

var AuthRoutes = func(router *mux.Router) {
	router.HandleFunc("/login", authcontrollers.Login).Methods("POST")
	router.HandleFunc("/register", authcontrollers.Register).Methods("POST")
	router.HandleFunc("/logout", authcontrollers.Logout).Methods("GET")
}
