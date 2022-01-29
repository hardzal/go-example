package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
)

func main() {
	// route /static/
	http.Handle("/assets/",
		http.StripPrefix("/assets/",
			http.FileServer(http.Dir("src\\basic-web\\sample\\3\\assets"))))
	// posisi running program menentukan directory diatas

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var filepath = path.Join("src\\basic-web\\sample\\3\\views", "index.html")
		var tmplt, err = template.ParseFiles(filepath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var data = map[string]interface{}{
			"title": "Learning Golang Web",
			"name":  "Migi",
		}

		err = tmplt.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("server started at localhost:8000")
	http.ListenAndServe(":8000", nil)
}
