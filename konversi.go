package main

import "fmt"

func main() {
	var a int32

	var b float32

	fmt.Print("Masukan angka : ")
	fmt.Scanf("%d", &a)
	b = float32(a)

	fmt.Printf("%.2f\n", b)

}
