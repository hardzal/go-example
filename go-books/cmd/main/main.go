package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hardzal/go-example/go-books/pkg/routes"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBooksStoreRoutes(r)
	routes.UsersStoresRoutes(r)
	http.Handle("/", r)
	fmt.Println("Starting web server at port :8001")
	log.Fatal(http.ListenAndServe("localhost:8001", r))
}
