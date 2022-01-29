package main

import "fmt"

func main()  {
	var nilai = 30

	// if nilai >= 8 && nilai <= 10{
	// 	fmt.Println("Lulus dengan hasil baik")
	// } else if nilai >= 6 && nilai <= 7 {
	// 	fmt.Println("Lulus dengan nilai rata-rata")
	// } else if nilai >= 0 && nilai <= 5{
	// 	fmt.Println("Kamu harus mengulang")
	// } else {
	// 	fmt.Println("Nilai kamu tidak terdefinisi")
	// }

	switch nilai {
		case 0:
			fmt.Println("it's not perfect")
		case 100:
			fmt.Println("it's perfect")
		case 1:
			fmt.Println("it's point")
		default:
		{
			fmt.Println("Average people")
			fmt.Println("You have chance to update that")
		}
	}

	var point = 2

	switch {
		case point == 10:
			fmt.Println("Perfect")
		case point >= 7 && point < 10:
			fmt.Println("Well done")
		case point >= 0 && point < 7:
			fmt.Println("Mengulang")
		default:
			{
				fmt.Println("it's over")
			}
	}

	// fallthrough : fungsinya seperti continue
	
}