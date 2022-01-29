package main

// setiap file program memiliki package
import "fmt"

func main() {
	var hello string
	hello = fmt.Sprintln("Hello, Word!")

	fmt.Println("Hello World!")
	fmt.Printf("%s", "Hello, World!")
	fmt.Print(hello)
}
