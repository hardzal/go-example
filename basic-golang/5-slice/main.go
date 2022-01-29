package main

import "fmt"

func main() {
	var nama []string
	var temp string
	var i = 0

	for i < 3 {
		fmt.Print("Nama : ")
		fmt.Scanln(&temp)
		nama = append(nama, temp)
		i++
	}

	fmt.Println(nama)

}
