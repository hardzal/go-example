package main

import (
	"fmt"
	"strings"
)

// Alias
// type FilterCallback func(string) bool

func filter(data []string, callback func(string) bool) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}

func main() {
	var data = []string{"wick", "jason", "ethan"}
	var dataContains = filter(data, func(each string) bool {
		return strings.Contains(each, "o")
	})
	var dataLength = filter(data, func(each string) bool {
		return len(each) == 5
	})

	fmt.Println("Data asli\t\t: ", data)
	fmt.Println("Filter ada horuf \"o\"\t: ", dataContains)
	fmt.Println("Filter jumlah huruf \"5\"\t : ", dataLength)
}
