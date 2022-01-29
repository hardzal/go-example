package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(5)

	// channel menghubungkan goroutine satu dengan yang lainnya, mekanismenya serah terima(estafet) data lewat channel tersebut.
	// pengirim dan penerima harus berada pada channel yang berbeda konsep ini disebut dengan buffered channel
	// pengiriman dan penerima channel bersifat blocking atau synchronous

	// deklarasi channel
	var messages = make(chan string)

	var sayHelloTo = func(who string) {
		var data = fmt.Sprintf("Hello %s", who)
		messages <- data
	}

	go sayHelloTo("John Cenna")
	go sayHelloTo("Ray Masterio")
	go sayHelloTo("Big Show")

	var message1 = <-messages
	fmt.Println(message1)

	var message2 = <-messages
	fmt.Println(message2)

	var message3 = <-messages
	fmt.Println(message3)
}
