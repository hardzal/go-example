package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// $ curl -X GET --user batman:secret http://localhost:9000/student
// $ curl -X GET --user batman:secret http://localhost:9000/student?id=s001
func main() {
	mux := http.DefaultServeMux

	mux.HandleFunc("/student", ActionStudent)

	var handler http.Handler = mux
	handler = MiddlewareAuth(handler)
	handler = MiddlewareAllowOnlyGET(handler)

	server := new(http.Server)
	server.Addr = ":8000"
	server.Handler = handler

	fmt.Println("server started at localhost:8000")
	server.ListenAndServe()
}

func ActionStudent(w http.ResponseWriter, r *http.Request) {
	if id := r.URL.Query().Get("id"); id != "" {
		OutputJSON(w, SelectStudent(id))
		return
	}

	OutputJSON(w, GetStudents())
}

func OutputJSON(w http.ResponseWriter, o interface{}) {
	res, err := json.Marshal(o)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	w.Write([]byte("\n"))
}
