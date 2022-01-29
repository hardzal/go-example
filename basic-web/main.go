package main

import (
	"fmt"
	"net/http"
)

// method index_handler handling index pages
func indexHandler(word http.ResponseWriter, res *http.Request) {
	fmt.Fprintf(word, "Whoaa, Golang!")
}

// method about_handler handling about pages
func aboutHandler(word http.ResponseWriter, res *http.Request) {
	fmt.Fprintf(word, "Expert golang developer!")
}

func contactHandler(word http.ResponseWriter, res *http.Request) {
	fmt.Fprintf(word, `
		<html>
			<head>
				<title>Hello, Golang!</title>
			</head>
			<body>	
	`)
	fmt.Fprintf(word, "<h1>Contact Us here</h1>")
	fmt.Fprintf(word, `
		<form method='post' action=''>
			<input type='text' name='nama'/>
	`)
	fmt.Fprintf(word, `<button type='submit'>Submit</button>
		</form>
	`)

	fmt.Fprintf(word, `
		</body>
		</html>
	`)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ohayou sekai o good morning world!"))
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/contact", contactHandler)
	http.HandleFunc("/hello", helloHandler)

	http.ListenAndServe(":8000", nil)
}
