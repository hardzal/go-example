package main

import (
	"fmt"
	"os"
)

// arguments adalah data opsional yang disisipkan ketika eksekusi program.
// flag merupakan ekstensi dari argument

// cara run
// go run ./main.go parameter here

func main() {
	// arguments
	var argsRaw = os.Args
	fmt.Printf("->%#v\n", argsRaw)

	var args = argsRaw[1:]
	fmt.Printf("->%#v\n", args)
}
