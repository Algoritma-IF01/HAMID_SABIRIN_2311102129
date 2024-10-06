package main

import "fmt"

func main() {
	var celsius int
	fmt.Print("Masukkan suhu dalam derajat Celsius: ")
	fmt.Scan(&celsius)

	fahrenheit := (celsius * 9 / 5) + 32
	reamur := celsius * 4 / 5
	kelvin := celsius + 273

	fmt.Printf("Suhu dalam Fahrenheit: %d\n", fahrenheit)
	fmt.Printf("Suhu dalam Reamur: %d\n", reamur)
	fmt.Printf("Suhu dalam Kelvin: %d\n", kelvin)
}
