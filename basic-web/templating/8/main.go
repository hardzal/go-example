package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const view string = `
	<html>
		<head>
			<title>Template</title>
		</head>
		<body>
			<h1>Hello, World!</h1>
		</body>
	</html>
`

func main() {
	http.HandleFunc("/view", func(w http.ResponseWriter, r *http.Request) {
		var tmpl = template.Must(template.New("index").ParseFiles("src\\basic-web\\templating\\view.html"))
		if err := tmpl.Execute(w, nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		var tmpl = template.Must(template.New("test").ParseFiles("src\\basic-web\\templating\\view.html"))
		if err := tmpl.Execute(w, nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		var tmpl = template.Must(template.New("main-template").Parse(view))
		if err := tmpl.Execute(w, nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/index", http.StatusTemporaryRedirect)
	})

	fmt.Println("server started at localhost:8000")
	http.ListenAndServe(":8000", nil)
}
