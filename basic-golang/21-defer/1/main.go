package main

import "fmt"

func foo() {
	defer fmt.Println("Done!")
	defer fmt.Println("Are we done?")
	fmt.Println("Let's doing something good")
}

func main() {
	foo()
}
