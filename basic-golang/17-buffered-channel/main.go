package main

import (
	"fmt"
	"runtime"
)

// channel secara default adalah un-buffered

func main() {
	runtime.GOMAXPROCS(2)

	messages := make(chan int, 2) // menentukan jumlah 'buffer' pada channel
	// jumlah buffer maksimal ada 3 (0, 1, 2)

	go func() {
		for {
			i := <-messages
			fmt.Println("Receive data", i)
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println("Send data", i)
		messages <- i
	}
}
