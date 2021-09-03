package main

import "fmt"

func main() {
	var celcius int

	fmt.Print("Masukan suhu dalam celcius : ")
	fmt.Scanf("%d", &celcius)
	kelvin := celcius + 273
	farenheit := (celcius * 9 / 5) + 32

	fmt.Println(celcius, "C = ", kelvin, "K")
	fmt.Println(celcius, "C =  ", farenheit, "F")
}
