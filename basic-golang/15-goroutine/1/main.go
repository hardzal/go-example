package main

import (
	"fmt"
	"runtime"
)

func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}

// concurent programming
// goroutine mirip bisa disebut juga sebagai 'mini thread'
func main() {
	// menentukan jumlah core yang diaktifkan saat eksekusi program
	runtime.GOMAXPROCS(5)

	// pembuatan goroutine baru
	go print(5, "How are you?")
	// print(5, "元気ですか？")
	print(5, "Shine!")

	var input string
	// menerima input baru
	fmt.Scanln(&input)
}
