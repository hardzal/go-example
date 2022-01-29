package main

import "fmt"

func main() {
	number := 10
	address := &number // memory address

	fmt.Println("Address of ", number, " is : ", address)
	fmt.Println("\nThe value of *address is : ", *address)

	*address = 5
	fmt.Println(number)
	fmt.Println(*address)

	number = number * *address
	fmt.Println(number)
	fmt.Println(*address)
}
