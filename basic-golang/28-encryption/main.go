package main

import (
	"crypto/sha1"
	"fmt"
	"time"
)

func doHashUsingSalt(text string) (string, string) {
	var salt = fmt.Sprintf("%d", time.Now().UnixNano())
	var saltedText = fmt.Sprintf("Text: '%s', salt: %s", text, salt)
	fmt.Println(saltedText)
	var sha = sha1.New()
	sha.Write([]byte(saltedText))
	var encrypted = sha.Sum(nil)

	return fmt.Sprintf("%x", encrypted), salt
}

func main() {
	var text = "secret"
	var sha = sha1.New()
	sha.Write([]byte(text))
	var encrypted = sha.Sum(nil)
	var encryptedString = fmt.Sprintf("%x", encrypted)
	fmt.Println("Original :", text)
	fmt.Println("Hashed :", encryptedString)

	// salt dalam kriptografi adalah data acak yang digabungkan pada data asli sebelum proses hash dilakukan
	// untuk menghindari hasil hash yang sama (dictionary attack)
	fmt.Println()
	fmt.Printf("Original : %s\n\n", text)
	var hashed1, salt1 = doHashUsingSalt(text)
	fmt.Printf("Hashed 1 : %s, salt: %s\n\n", hashed1, salt1)

	var hashed2, salt2 = doHashUsingSalt(text)
	fmt.Printf("Hashed 2 : %s, salt: %s\n\n", hashed2, salt2)

	var hashed3, salt3 = doHashUsingSalt(text)
	fmt.Printf("hashed 3 : %s, salt: %s\n\n", hashed3, salt3)
}
