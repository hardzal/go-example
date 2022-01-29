package main

import (
	"fmt"
	"math"
)

type hitung interface {
	luas() float64
	keliling() float64
}

type lingkaran struct {
	diameter float64
}

func (l lingkaran) jariJari() float64 {
	return l.diameter / 2
}

func (l lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jariJari(), 2)
}

func (l lingkaran) keliling() float64 {
	return math.Pi * l.diameter
}

type persegi struct {
	sisi float64
}

func (p persegi) luas() float64 {
	return math.Pow(p.sisi, 2)
}

func (p persegi) keliling() float64 {
	return p.sisi * 4
}

func main() {
	var bangundatar hitung

	bangundatar = persegi{10.5}
	fmt.Println("---Persegi---")
	fmt.Println("Luas\t\t: ", bangundatar.luas())
	fmt.Println("Keliling\t: ", bangundatar.keliling())

	bangundatar = lingkaran{21.0}
	fmt.Println("\n---Lingkaran---")
	fmt.Printf("luas\t\t: %.2f\n", bangundatar.luas())
	fmt.Printf("keliling\t: %.2f\n", bangundatar.keliling())
	fmt.Printf("jari - jari\t: %.2f\n", bangundatar.(lingkaran).jariJari())
}
