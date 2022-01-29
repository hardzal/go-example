package main

import (
	"fmt"
	"time"
)

func main() {
	var timeNow = time.Now()

	fmt.Printf("time: %v\n", timeNow)

	var timeDate = time.Date(2011, 12, 24, 10, 20, 0, 0, time.UTC)
	// time.Local
	fmt.Printf("timeDate: %v\n", timeDate)

}
