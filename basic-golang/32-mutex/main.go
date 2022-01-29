package main

import (
	"fmt"
	"runtime"
	"sync"
)

type counter struct {
	sync.Mutex
	val int
}

func (c *counter) Add(x int) {
	// mengeksklusifkan statement setelahnya
	// menjadikan hanya satu goroutine saja yang berjalan pada saat statement dijalankan
	c.Lock()
	c.val++
	c.Unlock()
}

func (c *counter) Value() (x int) {
	return c.val
}

func main() {
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup
	var meter counter
	// race condition
	// gunakan  flag -race untuk mendeteksi race condition
	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			for j := 0; j < 1000; j++ {
				meter.Add(1)
			}

			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(meter.Value())
}
