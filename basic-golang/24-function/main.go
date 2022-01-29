package main

import (
	"fmt"
	"strings"
)

func main() {
	var isExists = strings.Contains("john kenederz", "ken")
	var isPrefix = strings.HasPrefix("Johnnn", "j")
	var isSuffix = strings.HasSuffix("WTF", "F")
	var countStr = strings.Count("Hello, Welcome to new era", "e")
	var index1 = strings.Index("Ethan hunt", "ha")

	fmt.Println(isExists)
	fmt.Println(isPrefix)
	fmt.Println(isSuffix)
	fmt.Printf("Total karakter `e` : %d\n", countStr)
	fmt.Printf("Posisi index karakter `ha` : %d\n", index1)

	var text = "Hello, World!"
	var find = "l"
	var replaceWith = "r"

	var newText1 = strings.Replace(text, find, replaceWith, -1)
	fmt.Println(newText1)

	var str = strings.Repeat("la", 5)
	fmt.Println(str)

	var string1 = strings.Split("Hello, World", "")
	fmt.Println(string1)

	var data = []string{"Sejatinya", "Manusia", "adalah", "Lemah"}
	var newData = strings.Join(data, " ")
	fmt.Println(newData)

	// strings.ToUpper
	// strings.ToLower
}
