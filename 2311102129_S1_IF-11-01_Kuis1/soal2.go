package main

import "fmt"

func hitungmakanan(makanan, rombongan int) int {
	if makanan&makanan <= 3 {
		return 10000 * rombongan
	} else if makanan >= 4 && makanan <= 6 {
		return 2500 * rombongan * makanan
	} else if makanan > 50 {
		return 100000 * rombongan
	}
	return 0
}

func main() {
	var rombongan, makanan, jumlahorang int
	var sisa bool

	fmt.Print("jumlah rombongan : ")
	fmt.Scan(&rombongan)
	fmt.Print("jumlah makanan : ")
	fmt.Scan(&makanan)
	fmt.Print("jumlah orang : ")
	fmt.Scan(&jumlahorang)
	fmt.Print("apakah makanan sisa : ")
	fmt.Scan(&sisa)
	if sisa {
		fmt.Print("total yang harus dibayar : ", hitungmakanan(makanan, rombongan)+hitungmakanan(makanan, rombongan))
	} else {
		fmt.Print("total yang harus dibayar : ", hitungmakanan(makanan, rombongan))
	}
}
