package main

import (
	"fmt"
	"runtime"
	"sync"
)

// sync.waitGroup
// digunakan untuk menunggu goroutine

// sync untuk sinkronisasi goroutine

func doPrint(wg *sync.WaitGroup, message string) {
	defer wg.Done()
	fmt.Println(message)
}

func main() {
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		var data = fmt.Sprintf("Data %d", i)

		wg.Add(1)
		go doPrint(&wg, data)
	}

	wg.Wait()
}
