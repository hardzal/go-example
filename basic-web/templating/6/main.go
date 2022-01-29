package main

import (
	"fmt"
	"net/http"
	"text/template"
)

// Superhero struct
type Superhero struct {
	Name    string
	Alias   string
	Friends []string
}

// SayHello func struct
func (s Superhero) SayHello(from string, message string) string {
	return fmt.Sprintf("%s said: \"%s\"", from, message)
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var person = Superhero{
			Name:    "Bruce Wayne",
			Alias:   "Batman",
			Friends: []string{"Superman", "Flash", "Green Lantern"},
		}

		var tmplt = template.Must(template.ParseFiles("src\\basic-web\\templating\\view.html"))
		if err := tmplt.Execute(w, person); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("server started at localhost:8000")
	http.ListenAndServe(":8000", nil)
}
