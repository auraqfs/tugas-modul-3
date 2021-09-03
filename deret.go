package main

import "fmt"

func main() {
	var jumlah int
	var deret int

	fmt.Print("Masukan jumlah deret : ")
	fmt.Scan(&jumlah)
	fmt.Println("1. Deret ganjil")
	fmt.Println("2. Deret genap")
	fmt.Print("Pilih Deret [1/2] : ")
	fmt.Scan(&deret)

	for i := 0; i < jumlah; i++ {

		if deret == 2 {
			fmt.Print(" ", i*2)
		} else if deret == 1 {
			fmt.Print(" ", (i*2)+1)
		} else {
			fmt.Print("Masukan opsi deret kembali")
		}

	}
}
