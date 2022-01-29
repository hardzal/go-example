package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var data = "John cena"

	var encodedString = base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("Encoded :", encodedString)

	var decodedByte, _ = base64.StdEncoding.DecodeString(encodedString)
	var decodedString = string(decodedByte)

	fmt.Println("Decoded :", decodedString)

	var encoded = make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(encoded, []byte(data))
	var encodedString1 = string(encoded)
	fmt.Println("Encoded :", encodedString1)

	var url = "https://mrizal.id"

	var encodedString2 = base64.URLEncoding.EncodeToString([]byte(url))
	fmt.Println("Url encoded :", encodedString2)

	var decodedByte1, _ = base64.URLEncoding.DecodeString(encodedString2)
	var decodedString1 = string(decodedByte1)
	fmt.Println("Url decoded :", decodedString1)

}
