package main

import "fmt"

func main() {
	grades := make(map[string]float32)

	grades["Luffy"] = 19
	grades["Zoro"] = 21

	fmt.Println(grades)

	Zoro := grades["Zoro"]
	fmt.Println(Zoro)

	delete(grades, "Zoro")
	fmt.Println(Zoro)
	fmt.Println(grades)
}
