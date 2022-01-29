package main

import (
	"fmt"
	"runtime"
)

func printMessage(what chan string) {
	fmt.Println(<-what)
}

func main() {
	runtime.GOMAXPROCS(5)

	var messages = make(chan string)

	for _, each := range []string{"Luffy", "Zoro", "Sanji"} {
		go func(who string) {
			var data = fmt.Sprintf("Hello %s", who)
			messages <- data
		}(each)
	}

	for i := 0; i < 3; i++ {
		printMessage(messages)
	}
}
