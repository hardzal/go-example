package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func cleanup() {
	defer wg.Done()
	if r := recover(); r != nil {
		fmt.Println("Recovered in cleanup", r)
	}
}

func say(s string) {
	defer cleanup()
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
		if i == 2 {
			panic("It's 2")
		}
	}
}

func main() {
	// akan dijalankan setelah program menyelesaikan tugas
	wg.Add(1)
	go say("Hey!!")
	wg.Add(1)
	go say("Hello!")
	wg.Wait()

	// tugas ini yang dimaksud
	// time.Sleep(time.Millisecond)
}
