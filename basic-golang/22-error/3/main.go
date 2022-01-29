package main

import (
	"errors"
	"fmt"
	"strings"
)

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("Cannot be empty")
	}
	return true, nil
}

func catch() {
	// IIFE
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error occured", r)
		} else {
			fmt.Println("Application running perfectly")
		}
	}()
}

func main() {
	var name string
	fmt.Print("Input your name : ")
	fmt.Scanln(&name)

	if valid, err := validate(name); valid {
		fmt.Println("Hello,", name)
	} else {
		fmt.Println(err.Error())
		panic(err.Error())
		// fmt.Println("")
	}
}
