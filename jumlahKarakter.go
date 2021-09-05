package main

import (
	"fmt"
)

func main() {
	var str string

	fmt.Print("Masukan karakter : ")
	fmt.Scan(&str)

	fmt.Println("Jumlah karakternya : ", len(str))
}
