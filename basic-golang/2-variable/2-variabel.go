package main

import "fmt"

func main() {
	var firstName string = "M"
	// manifest typing: tipe data yang digunakan harus dituliskan juga
	var lastName string
	lastName = "Rizal"
	// deklarasi sekaligus assignment
	jenisKelamin := "laki - laki"

	fmt.Printf("Hello, %s %s\n", firstName, lastName)
	fmt.Printf("Welcome back " + firstName + " " + lastName + "\n")
	fmt.Println(firstName, jenisKelamin)
	// komentar singleline
	/*
		multiline
	*/
	var first, second, third string
	first, second, third = "satu", "dua", "tiga"
	fourth, fifth, sixth := "empat", "lima", "enam"
	fmt.Println(first, second, third, fourth, fifth, sixth)
	// unused variabel
	// predefined variabel
	_ = "belajar golang mudah"

	name := new(string)
	fmt.Println(name)
	fmt.Println(*name)
	// deklarasi variabel menggunakan keyword make
	// channel
	// slice
	// map

	// uint8, uint16, uint32, uint65, uint, byte, int8, int16, int32, int64, int, rune
	// float32, float64
	// bool
	// string
	// complex128

	// nil
	// pointer
	// slice
	// map
	// channel
	// interface()

	// const

	// %s (string), %d (integer), %f (float), %t (bool)
	// %v semua jenis data

}
