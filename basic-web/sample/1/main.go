package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var message = "おはよう　せかい、Good Morning World!"
	w.Write([]byte(message))
}

func main() {
	http.HandleFunc("/", indexHandler)

	var host = "localhost:8000"
	fmt.Printf("Server started at %s\n", host)
	server := new(http.Server)
	server.Addr = host
	err := server.ListenAndServe()

	if err != nil {
		fmt.Println(err.Error())
	}
}
