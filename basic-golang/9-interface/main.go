package main

import "fmt"

func main() {
	// type data interface
	// var secret interface{}

	// secret = "ethan hunt"
	// fmt.Println(secret)

	// secret = []string{"apple", "manggo", "banana"}
	// fmt.Println(secret)

	// secret = 12.4
	// fmt.Println(secret)

	var data map[string]interface{}

	data = map[string]interface{}{
		"name":   "John wick",
		"grade":  9,
		"dinner": []string{"Fried chicken", "hot tea", "coffee"},
	}

	fmt.Println(data)
}
