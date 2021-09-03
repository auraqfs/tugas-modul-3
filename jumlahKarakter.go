package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	var jumlah = strings.Count(str, "a")

	fmt.Print("Masukan karakter : ")
	fmt.Scan(&str)

	fmt.Println("Jumlah karakter yaitu ", jumlah)
}
