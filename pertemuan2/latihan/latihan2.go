package main

import "fmt"

func main() {
	var tahun int

	fmt.Print("Masukkan tahun: ")

	fmt.Scanln(&tahun)

	if tahun%4 == 0 {
		fmt.Println("Kabisat: true")
	} else {
		fmt.Println("Kabisat: false")
	}
}
