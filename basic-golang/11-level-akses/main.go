package main

import (
	"basic-golang/11-level-akses/library"
	"fmt"
)

func main() {
	library.SayHello("Nisfu")

	// tidak bisa mengakses .introduce karena memiliki level akses private
	// library.introduce("nama")
	// murid
	var murid = library.StudentCollege{"M Rizal", "123170037"}

	fmt.Println(murid.Name)
}
