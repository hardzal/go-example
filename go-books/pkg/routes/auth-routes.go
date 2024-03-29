package routes

import (
	"github.com/gorilla/mux"
	"github.com/hardzal/go-example/go-books/pkg/controllers/authcontrollers"
	"github.com/hardzal/go-example/go-books/pkg/controllers/productcontroller"
	"github.com/hardzal/go-example/go-books/pkg/middlewares"
)

var AuthRoutes = func(router *mux.Router) {
	router.HandleFunc("/login", authcontrollers.Login).Methods("POST")
	router.HandleFunc("/register", authcontrollers.Register).Methods("POST")
	router.HandleFunc("/logout", authcontrollers.Logout).Methods("GET")
	router.HandleFunc("/logout", authcontrollers.Logout).Methods("GET")

	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/products", productcontroller.Index).Methods("GET")
	api.Use(middlewares.JWTMiddleware)
}
