package main

import (
	"fmt"
	"reflect"
)

// Reflect
// mengambil informasi dari variabel tersebut

type student struct {
	Name  string
	Grade int
}

func (s *student) getPropertyInfo() {
	var reflectValue = reflect.ValueOf(s)

	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}

	var reflectType = reflectValue.Type()

	for i := 0; i < reflectValue.NumField(); i++ {
		fmt.Println("Nama\t\t: ", reflectType.Field(i).Name)
		fmt.Println("Tipedata\t: ", reflectType.Field(i).Type)
		fmt.Println("Nilai\t\t: ", reflectValue.Field(i).Interface())
		fmt.Println("")
	}
}

func (s *student) SetName(name string) {
	s.Name = name
}

func main() {
	var number = 23
	var reflectValue = reflect.ValueOf(number)
	var mahasiswa = &student{Name: "Johny", Grade: 10}

	fmt.Println("Tipe variabel\t: ", reflectValue.Type())
	fmt.Println("Nilai Variabel\t: ", reflectValue.Interface())
	if reflectValue.Kind() == reflect.Int {
		fmt.Println("Nilai variabel\t: ", reflectValue.Int())
	}
	fmt.Println("")
	mahasiswa.getPropertyInfo()

	reflectValue = reflect.ValueOf(mahasiswa)
	var method = reflectValue.MethodByName("SetName")
	method.Call([]reflect.Value{
		reflect.ValueOf("Johny"),
	})

	fmt.Println("Nama\t: ", mahasiswa.Name)
}
