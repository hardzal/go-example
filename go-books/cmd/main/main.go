package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hardzal/go-example/go-books/routes"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBooksStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8001", r))
}
