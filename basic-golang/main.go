package main

import (
	"fmt"
	"time"
)

func main() {
	independenceDay := time.Date(1945, 8, 17, 10, 0, 0, 0, time.Local)
	const country = "Republik Indonesia"
	var age = time.Now().Year() - independenceDay.Year()

	fmt.Printf("Dirgahayu %s ke-%d\n", country, age)
}
