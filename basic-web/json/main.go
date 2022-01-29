package main

import (
	"encoding/json"
	"fmt"
)

// User struct
type User struct {
	FullName string `json:"Name"`
	Age      int
}

func main() {
	var jsonString = `[
		{"Name": "John Cenna", "Age": 25 },
		{"Name": "Senku", "Age": 16}
	]`
	var jsonData = []byte(jsonString)

	var data []User
	// var data1 map[string]interface{}

	// menggunakan interface{}
	// tetapi harus terlebih dahulu dicasting
	// var dat2decoded = data2.(map[string]interface{})

	var err = json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// json.Unmarshal(jsonData, &data1)

	for i := 0; i < len(data); i++ {
		fmt.Println("User\t: ", data[i].FullName)
		fmt.Println("Age\t: ", data[i].Age)
		fmt.Println()
	}

	// fmt.Println("User\t: ", data1["Name"])
	// fmt.Println("Age\t: ", data1["Age"])
}
