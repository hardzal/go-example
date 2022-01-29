package main

import (
	"basic-golang/10-package/person"
	"fmt"
)

func main() {
	p := person.Description("Momo")
	fmt.Println(p)
}
