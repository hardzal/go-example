package library

import "fmt"

// struct juga memiliki level akses public dan private
var Student = struct {
	Name  string
	Grade int
}{}

func init() {
	Student.Name = "John Wick"
	Student.Grade = 10

	fmt.Println("-->library/Student.go exported")
}
