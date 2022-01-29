package main

import "fmt"

type mataKuliah struct {
	code int
	rate float32
}

type student struct {
	name  string
	nim   int
	rates mataKuliah
}

type person struct {
	name   string
	age    int
	gender string
}

func main() {
	A := person{name: "John", age: 19, gender: "Pria"}
	fmt.Println("Name\t:", A.name)
	fmt.Println("Age\t:", A.age)
	fmt.Println("Gender\t:", A.gender)

	B := &person{name: "Lala", age: 17, gender: "Wanita"}
	fmt.Println("Name\t", B.name)
	fmt.Println("Name\t", A.name)

	A.show()
	B.show()
	B.setName("Momo")
	B.setAge(20)
	B.show()
	A.setName("Nana")
	A.setAge(16)
	A.show()
	fmt.Println(*B, A)
}

func (p *person) show() {
	fmt.Printf("%v is %v years old.\n", p.name, p.age)
}

func (p *person) setAge(age int) {
	p.age = age
}

func (p person) setName(name string) {
	p.name = name
}
