package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str = "1234"
	// convertion string to int
	var num, err = strconv.Atoi(str)
	// convertion int to string
	// strconv.Itoa()
	//
	if err == nil {
		num = 5 + num
		str = fmt.Sprintf("Hasilnya : %d", num)
		fmt.Printf("%s\n", str)
	}

	// convertion dari string ke int dengan tipe int yang bisa ditentukan
	var str1 = "4321"
	var str2 = "10101"
	// basis 10 dengan tipe data int64
	var num1, err1 = strconv.ParseInt(str1, 10, 64)
	var num2, err2 = strconv.ParseInt(str2, 2, 8)
	if err1 == nil {
		fmt.Printf("%d\n", num1)
	}
	if err2 == nil {
		fmt.Printf("%d\n", num2)
	}

	var num3 = int64(30)
	var str3 = strconv.FormatInt(num3, 8)
	fmt.Printf("Konversi basis 10 ke basis 8 : %d -> %s\n", num3, str3)

	var num4 = float64(3.1412411512131)
	var str4 = strconv.FormatFloat(num4, 'f', 4, 64)
	fmt.Println(str4)

	var str5 = "true"
	var bol, err3 = strconv.ParseBool(str5)

	if err3 == nil {
		fmt.Println(bol == false)
	}

	var str6 = strconv.FormatBool(bol)
	fmt.Printf("%s\n", str6)

	var data = map[string]interface{}{
		"nama":    "John",
		"grade":   5,
		"height":  156.5,
		"isMale":  true,
		"hobbies": []string{"eating", "sleeping"},
	}

	// fmt.Println(data["nama"].(string))
	// fmt.Println(data["grade"].(int))
	// fmt.Println(data["height"].(float64))
	// fmt.Println(data["isMale"].(bool))
	// fmt.Println(data["hobbies"].([]string))

	for _, val := range data {
		switch val.(type) {
		case string:
			fmt.Println(val.(string))
		case int:
			fmt.Println(val.(int))
		case float64:
			fmt.Println(val.(float64))
		case bool:
			fmt.Println(val.(bool))
		case []string:
			fmt.Println(val.([]string))
		default:
			fmt.Println(val)
		}
	}
}
