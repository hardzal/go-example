package main

import (
	"fmt"
	"regexp"
)

func main() {

	var text = "banana,burger,soup"
	var regex, err = regexp.Compile(`[a-z]+`)

	if err != nil {
		fmt.Println(err.Error())
	}

	var res1 = regex.FindAllString(text, 2)
	fmt.Printf("%v\n", res1)

	var isMatch = regex.MatchString(text)
	fmt.Println(isMatch)

	// regex.FindString() // menghasilkan string yang sesuai pattern regex
	// regex.FindStringIndex() // mirip dengan FindString() hanya saja mengembalikan nilai index

	var replace = regex.ReplaceAllString(text, "vegetables")
	fmt.Println(replace)

	var str = regex.ReplaceAllStringFunc(text, func(each string) string {
		if each == "burger" {
			return "potato"
		}
		return each
	})

	fmt.Println(str)

	var splitRegex = regex.Split(text, -1)
	fmt.Printf("%v\n", splitRegex)
}
