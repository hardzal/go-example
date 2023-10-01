package main

import (
	"log"
	"net/http"

	"github.com/hardzal/go-example/mux-mongo-api/configs"
	"github.com/hardzal/go-example/mux-mongo-api/routes"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	configs.ConnectDB()

	// routes
	routes.UserRoute(router)

	// router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
	// 	rw.Header().Set("Content-Type", "application/json")

	// 	json.NewEncoder(rw).Encode(map[string]string{"data": "Hello From Mux & MongoDB"})
	// }).Methods("GET")

	log.Fatal(http.ListenAndServe(":6000", router))
}
