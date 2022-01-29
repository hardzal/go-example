package main

import (
	"fmt"
	"io"
	"os"
)

var path = "./src/basic-golang/31-file/new-file.txt"

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

func createFile() {
	var _, err = os.Stat(path)

	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err) {
			return
		}
		defer file.Close()
	}

	fmt.Println("==> file berhasil dibuat", path)
}

func writeFile() {
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	_, err = file.WriteString("Hello, World!\n")
	if isError(err) {
		return
	}
	_, err = file.WriteString("Let's learn Golang\n")
	if isError(err) {
		return
	}

	err = file.Sync()
	if isError(err) {
		return
	}

	fmt.Println("-->File berhasil diperbaharui")
}

func readFile() {
	var file, err = os.OpenFile(path, os.O_RDONLY, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	var text = make([]byte, 1024)
	for {
		n, err := file.Read(text)
		if err != io.EOF {
			if isError(err) {
				break
			}
		}
		if n == 0 {
			break
		}
	}

	if isError(err) {
		return
	}

	fmt.Printf("==>Mulai membaca isi file...\n")
	fmt.Println(string(text))
}

func deleteFile() {
	var err = os.Remove(path)
	if isError(err) {
		return
	}

	fmt.Println("File berhasil dihapus!")
}

func main() {
	// membuat file
	// createFile()

	// menulis file
	// writeFile()

	// membaca file
	// readFile()

	// menghapus file
	// deleteFile()
}
