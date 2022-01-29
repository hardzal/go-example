// Actions & Variables
package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// Info struct
type Info struct {
	Affliation string
	Address    string
}

// Person struct
type Person struct {
	Name    string
	Gender  string
	Hobbies []string
	Info    Info
}

// GetAffliationDetailInfo func struct
func (t Info) GetAffliationDetailInfo() string {
	return "have 27 divisions"
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var person = Person{
			Name:    "Bruce Wayne",
			Gender:  "male",
			Hobbies: []string{"Reading Books", "Travelling", "Buying things"},
			Info:    Info{"Wayne Enterprises", "Gotham City"},
		}

		var tmplt = template.Must(template.ParseFiles("src\\basic-web\\templating\\view.html"))
		if err := tmplt.Execute(w, person); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("server started at localhost:8000")
	http.ListenAndServe(":8000", nil)
}
