package main

import (
	"fmt"
	"runtime"
)

// for - range - close untuk handle penerimaan data

func sendMessage(ch chan<- string) {
	for i := 0; i < 20; i++ {
		ch <- fmt.Sprintf("Data : %d", i)
	}
	close(ch)
}

func printMessage(ch <-chan string) {
	for message := range ch {
		fmt.Println(message)
	}
}

func main() {
	runtime.GOMAXPROCS(2)

	var messages = make(chan string)
	go sendMessage(messages)
	printMessage(messages)
}

// channel direction
// ch chan string // digunakan untuk mengirim dan menerima data
// ch chan<- string // digunakan untuk mengirim data
// ch<- chan string // digunakan untuk menerima data
