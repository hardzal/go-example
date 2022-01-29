package main

import (
	"flag"
	"fmt"
)

func main() {
	// flags
	var name = flag.String("name", "anonymous", "what's name?")
	var age = flag.Int64("age", 22, "type your age")
	var status string
	var gender bool
	var sex string

	flag.StringVar(&status, "status", "single", "Relationship status")
	flag.BoolVar(&gender, "gender", true, "Gender")

	flag.Parse()

	if gender == true {
		sex = "Male"
	} else {
		sex = "Female"
	}

	fmt.Printf("Name\t: %s\n", *name)
	fmt.Printf("Age\t: %d\n", *age)
	fmt.Printf("Gender\t: %s\n", sex)
	fmt.Printf("Status\t: %s\n", status)
}
