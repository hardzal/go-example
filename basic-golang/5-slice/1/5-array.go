package main

import "fmt"

func main() {
	var names [5]string
	names[0] = "hello"
	names[1] = "world"
	names[2] = "new"
	names[3] = "world"
	names[4] = "new order"

	fmt.Println(names[0], names[1], names[2], names[3], names[4])

	var hitung = [3]uint{1, 2, 3}
	var hobby = [3]string{
		"membaca",
		"menulis",
		"menonton",
	}
	for i := 0; i < len(hitung); i++ {
		fmt.Println(hitung[i])
	}
	var i = 0
	for i < len(hobby) {
		fmt.Println("hobbyku :", hobby[i])
		i++
	}

	var matriks1 = [3][3]int{[3]int{1, 2, 3}, [3]int{4, 5, 6}, [3]int{7, 8, 9}}
	fmt.Println("Matriks :", matriks1)
	for _, matrik := range matriks1 {
		fmt.Println(matrik)
	}

	// alokasi menggunakan 'make'
	var bahasa = make([]string, 3)
	bahasa[0] = "Indonesia"
	bahasa[1] = "Inggris"
	bahasa[2] = "Jepang"
	fmt.Println(bahasa)
	// len() cap()
	// slice
	var fruits = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(fruits[2])
	// tipe data reference
	var newFruits = fruits[0:3]
	// [3:] [:3] [:]
	fmt.Println(newFruits)
	newFruits = append(newFruits, "manggo")
	fmt.Println(newFruits)

	dst := []string{"potate", "manggo", "grape"}
	src := []string{"watermelon"}
	n := copy(dst, src)
	fmt.Println(n)
	fmt.Println(dst)

	var nilai = [...]int{1, 2, 3, 4, 5}
	fmt.Println("Value of nilai: ", nilai)
	fmt.Println("Banyaknya isi nilai: ", len(nilai))
}
