package main

import (
	"fmt"
)

func main() {
	fmt.Println("Masukkan 5 angka integer antara 32 sampai 127:")

	var angka1, angka2, angka3, angka4, angka5 int
	fmt.Print("masukan angka: ")
	fmt.Scan(&angka1, &angka2, &angka3, &angka4, &angka5)

	fmt.Println("Masukkan beberapa karakter dengan angka ASCII:")

	var input string
	fmt.Scan(&input)

	// Output
	fmt.Println("\nOutput:")

	fmt.Println("Karakter dari angka-angka yang dimasukkan:", angka1, angka2, angka3, angka4, angka5)

	fmt.Printf("%c%c%c%c%c\n", angka1, angka2, angka3, angka4, angka5)

	if len(input) > 0 {
		fmt.Println("Karakter yang diinkrementasi dengan +1:")
		for i := 0; i < len(input); i++ {
			fmt.Printf("%c", input[i]+1)
		}
		fmt.Println()
	}
}

