package main

import "fmt"

func main() {
	var greeting = "Selamat datang di dap"
	var a, b int

	fmt.Println(greeting)

	fmt.Print("Masukkan dua angka: ")
	fmt.Scanln(&a, &b)

	fmt.Printf("%d + %d = %d\n", a, b, a+b)
}
