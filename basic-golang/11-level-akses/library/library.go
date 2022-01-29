package library

import "fmt"

// Student struct
type StudentCollege struct {
	Name string
	Nim  string
}

// SayHello memiliki level akses public karena menggunakan Capital diawal
func SayHello(name string) {
	fmt.Println("Hello!")
	introduce(name)
}

// method intrdoce memiliki level akses private
func introduce(name string) {
	fmt.Println("Hello, nama saya", name)
}
