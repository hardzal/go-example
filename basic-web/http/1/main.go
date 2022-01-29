package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			w.Write([]byte("POST"))
		} else if r.Method == "GET" {
			w.Write([]byte("GET"))
		} else {
			http.Error(w, "", http.StatusBadRequest)
		}
	})

	fmt.Println("server started at localhost:8000")
	http.ListenAndServe(":8000", nil)
}
