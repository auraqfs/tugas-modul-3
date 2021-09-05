package main

import (
	"fmt"

	"github.com/yudapc/go-rupiah"
)

func main() {
	var amount int

	fmt.Print("Masukan nilai : ")
	fmt.Scanf("%d", &amount)
	amountFloat := float64(amount)
	formatRupiah := rupiah.FormatRupiah(amountFloat)
	fmt.Println(formatRupiah)

	formatDollar := amountFloat / 14250
	fmt.Printf("$ %.2f", formatDollar)
}
