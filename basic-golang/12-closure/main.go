package main

import "fmt"

// closure adalah sebuah fungsi yang disimpan dalam sebuah variabel

func findMax(numbers []int, max int) (int, func() []int) {
	var res []int
	for _, e := range numbers {
		if e <= max {
			res = append(res, e)
		}
	}
	return len(res), func() []int {
		return res
	}
}

func main() {
	var numbers = []int{3, 1, 2, 5, 9, 10, 4, 5, 12, 8, 6, 7}
	var filter = 5
	var getMinMax = func(n []int) (int, int) {
		var min, max int
		for i, x := range n {
			switch {
			case i == 0:
				max, min = x, x
			case x > max:
				max = x
			case x < min:
				min = x
			}
		}
		return min, max
	}

	var min, max = getMinMax(numbers)
	fmt.Printf("Data : %v\nMin : %v\nMax : %v\n\n", numbers, min, max)

	// IIFE
	var newNumbers = func(min int) []int {
		var r []int
		for _, e := range numbers {
			if e < min {
				continue
			}
			r = append(r, e)
		}
		return r
	}(filter)

	fmt.Println("original Number : ", numbers)
	fmt.Println("filtered Number : ", newNumbers)

	var howMany, getNumbers = findMax(numbers, filter)
	var theNumbers = getNumbers()

	fmt.Println("Numbers\t: ", numbers)
	fmt.Printf("filter\t : %d\n\n", filter)

	fmt.Println("Found\t : ", howMany)
	fmt.Println("Value\t :", theNumbers)
}
