package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var data = map[string]string{
			"name":    "John Wick",
			"Message": "Have a nice day",
		}

		var t, err = template.ParseFiles("src\\basic-web\\templating\\2\\index.html")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		t.Execute(w, data)
	})

	fmt.Println("Starting web server at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
