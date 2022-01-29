package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type M map[string]interface{}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// var tmpltAll, err = template.ParseGlob("src\\basic-web\\sample\\views\\*")
		// if err != nil {
		// 	panic(err.Error())
		// 	return
		// }

		var tmpltAll = template.Must(template.ParseFiles(
			"src\\basic-web\\sample\\views/index.html",
			"src\\basic-web\\sample\\views/_header.html",
			"src\\basic-web\\sample\\views/_message.html",
		))

		var data = M{"name": "Ningen"}
		var err = tmpltAll.ExecuteTemplate(w, "index", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		// var tmpltAll, err = template.ParseGlob("src\\basic-web\\sample\\views\\*")
		// if err != nil {
		// 	panic(err.Error())
		// 	return
		// }

		var tmpltAll = template.Must(template.ParseFiles(
			"src\\basic-web\\sample\\views/about.html",
			"src\\basic-web\\sample\\views/_header.html",
			"src\\basic-web\\sample\\views/_message.html",
		))

		var data = M{"name": "Hito"}
		var err = tmpltAll.ExecuteTemplate(w, "about", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("server started at localhost:8000")
	http.ListenAndServe(":8000", nil)
}
