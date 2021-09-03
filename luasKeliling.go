package main

import "fmt"

func main() {
	var opsi, p, l int

	fmt.Println("Pilih Opsi")
	fmt.Println("1. Hitung Luas")
	fmt.Println("2. Hitung Keliling")
	fmt.Print("Masukan opsi : ")
	fmt.Scanln(&opsi)
	fmt.Println("=====================================================")

	if opsi == 1 {
		fmt.Print("Masukan Panjang Persegi Panjang : ")
		fmt.Scan(&p)

		fmt.Print("Masukan Lebar Persegi Panjang : ")
		fmt.Scan(&l)
		luas := p * l

		fmt.Println("Luas Persegi Panjang yaitu ", luas)
	} else if opsi == 2 {
		fmt.Print("Masukan Panjang Persegi Panjang : ")
		fmt.Scan(&p)

		fmt.Print("Masukan Lebar Persegi Panjang : ")
		fmt.Scan(&l)
		keliling := (p + l) * 2

		fmt.Println("Keliling Persegi Panjang yaitu ", keliling)
	} else {
		fmt.Print("Masukan opsi kembali")
	}

}
