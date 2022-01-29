package main

import (
	"fmt"

	"github.com/novalagung/gubrak"
)

// dep init
// dep ensure -update github.com/novalagung/gubrak

func main() {
	fmt.Println(gubrak.RandomInt(10, 20))
}
