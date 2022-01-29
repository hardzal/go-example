package main

import "fmt"

type person struct {
	name string
	age  int
}

type student struct {
	name        string
	height      float64
	age         int32
	isGraduated bool
	hobbies     []string
}

func main() {
	// anonymous struct
	var c = struct {
		person
		grade int
	}{}

	var data = student{
		name:        "John",
		height:      188.5,
		age:         23,
		isGraduated: true,
		hobbies:     []string{"Swimming", "Running"},
	}

	c.person = person{"John", 23}
	c.grade = 8
	fmt.Println(c)

	// %b data numerik menjadi biner
	fmt.Printf("%b\n", data.age)
	// %c data numerik menjadi kode unicodenya
	fmt.Printf("%c\n", data.age)
	// %e atau %E data numerik desimal ke bentuk notasi numerik standar Scientific notation
	fmt.Printf("%e\t%E\n", data.height, data.height)
	// %f atau %F data numerik menjadi bilangan desimal yang
	fmt.Printf("%.2f\n", data.height)
	// %g atau %G untuk menformat data numerik menjadi desimal sesuai panjang nilai desimalnya
	fmt.Printf("%g\n", 0.123123123123)
	// %o menformat data numerik menjadi bentuk string numerik berbasis oktal
	fmt.Printf("%o\n", data.age)
	// %p menformat data pointer, mengembalikan alamat pointer referensi variabel
	fmt.Printf("%p\n", &data.name)
	// %q untuk escape string
	fmt.Printf("%q\n", `" name \ height "`)
	fmt.Printf("%s\n", data.name)
	// memformat data boolean
	fmt.Printf("%t\n", data.isGraduated)

	// %T Mengambil tipe variabelnya
	fmt.Printf("%T\n", data.name)

	// %v digunakan untuk segala jenis tipe data
	fmt.Printf("%v\n", data)

	// %+v digunakan untuk memformat tipe data struct sesuai urutan property
	fmt.Printf("%+v\n", data)

	// %#v digunakan untuk memformat tipe data struct dan menampilkan bagaimana struct dideklarasikan
	fmt.Printf("%#v\n", data)

	// %x untuk memformat data numerik menjadi bilangan hexadesimal
	fmt.Printf("%x\n", data.age)

	// %% menulis karakter %
}
