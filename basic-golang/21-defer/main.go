package main

import (
	"fmt"
)

func orderSomeFood(menu string) {
	defer fmt.Println("Terimakasih, silahkan menunggu!")
	if menu == "Pizza" {
		fmt.Print("Pilihan tepat!", " ")
		fmt.Print("Pizza ditempat kami paling enak!", "\n")
		return
	}

	fmt.Println("Pesanan Anda : ", menu)
}

func main() {
	orderSomeFood("Pizza")
	// os.Exit(1)
	orderSomeFood("Burger")
}
