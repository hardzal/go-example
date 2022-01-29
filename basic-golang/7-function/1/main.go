package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

func main() {
	var names = []string{"John", "Sin"}
	printMessage("Hello", names)

	rand.Seed(time.Now().Unix())
	var randomValue int
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number : ", randomValue)

	divideNumber(5, 2)
	var area, circumference = calculateCircle(14)
	fmt.Println("Luas lingkaran\t\t= ", area)
	fmt.Println("Keliling Lingkaran\t= ", circumference)

	var volumCube, areaCube = calculateCube(8)
	fmt.Printf("Luas Kubus\t: %.1f\n", areaCube)
	fmt.Printf("Volume Kubus\t: %.1f\n", volumCube)
	fmt.Println()

	var numbers = sumNumber(3, 4, 5, 6, 7, 8, 9)

	fmt.Println("Total ", numbers)

	var angka = []int64{5, 6, 1, 9, 10}
	var total = sumNumber(angka...)
	fmt.Println("Total angka\t= ", total)

	i := 0
	increment(&i)
	fmt.Println("Hasil penambahan nilai =", i)
}

func printMessage(message string, arr []string) {
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)
}

func randomWithRange(min int, max int) int {
	var value = rand.Int()%(max-min+1) + min
	return value
}

func divideNumber(m, n float64) {
	if n == 0 {
		fmt.Printf("Invalid divider. %.2f cannot divided by %.2f\n", m, n)
		return
	}

	var result = m / n
	fmt.Printf("%.2f / %.2f = %.2f\n", m, n, result)
}

func calculateCircle(d float64) (float64, float64) {
	var area = math.Pi * math.Pow(d/2, 2)
	var circumference = math.Pi * d

	return area, circumference
}

func calculateCube(s float64) (volum, area float64) {
	area = math.Pow(s, 3)
	volum = 6 * s * s
	return
}

// function variadic
func sumNumber(numbers ...int64) int64 {
	var total int64 = 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func yourHobbies(name string, hobbies ...string) {
	var hobbiesAsString = strings.Join(hobbies, ", ")

	fmt.Printf("Hello, my name is %s\n", name)
	fmt.Printf("My hobbies are: %s\n", hobbiesAsString)
}

func increment(i *int) {
	*i++
}
