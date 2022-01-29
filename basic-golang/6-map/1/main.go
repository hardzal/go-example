package main

import "fmt"

func main() {
	var bulan = map[string]int{"januari": 31, "februari": 28, "maret": 31}
	// fmt.Println("januari", bulan["januari"])
	for key, val := range bulan {
		fmt.Println(key, " \t:", val)
	}
	delete(bulan, "februari")
	if bulan["februari"] != 0 {
		fmt.Println(bulan["februari"])
	} else {
		fmt.Println("data bulan februari tidak ada")
		fmt.Println(len(bulan))
	}

	var language = map[string]int{"html": 70, "css": 60, "php": 50, "python": 10, "ruby": 25}

	// menghapus data amap
	delete(language, "ruby")
	var value, isExist = language["ruby"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exists")
	}

	var dataMahasiswa = []map[string]string{
		map[string]string{"name": "M Rizal", "gender": "laki-laki"},
		map[string]string{"name": "Bangkit kit", "gender": "lelaki"},
	}

	var _ = []map[string]string{
		{"name": "Angku", "gender": "other", "birth place": "Raftel"},
		{"address": "Jalan grand line"},
	}

	for _, mahasiswa := range dataMahasiswa {
		fmt.Println(mahasiswa["name"], "\t: ", mahasiswa["gender"])
	}
}
