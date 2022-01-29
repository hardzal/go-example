package main

import "fmt"

// kumpulan method yang tidak memiliki isi hanya deklarasi

type animal interface {
	description() string
}

type cat struct {
	Type  string
	Sound string
}

type snake struct {
	Type      string
	Poisonous bool
}

func (s snake) description() string {
	return fmt.Sprintf("Poisonous\t: %v", s.Poisonous)
}

func (c cat) description() string {
	return fmt.Sprintf("Sound\t\t: %v", c.Sound)
}

func main() {
	var a animal
	a = snake{Poisonous: true}
	fmt.Println(a.description())
	a = cat{Sound: "Meowuu!!"}
	fmt.Println(a.description())
}
