package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func foo(c chan int, someValue int) {
	defer wg.Done()
	c <- someValue * 5
}

func main() {
	// channel buffer
	fooVal := make(chan int, 10)

	// go foo(fooVal, 5)
	// go foo(fooVal, 2)
	// go foo(fooVal, 3)

	// v1 := <-fooVal
	// v2 := <-fooVal
	// v3 := <-fooVal

	// fmt.Println(v1)
	// fmt.Println(v2)
	// fmt.Println(v3)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go foo(fooVal, i)
	}

	wg.Wait()
	close(fooVal)

	for item := range fooVal {
		fmt.Println(item)
	}
}
