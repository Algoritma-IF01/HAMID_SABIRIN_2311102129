package main

import (
	"fmt"
)

func angkalipatan(arr []int, index int) int {
	if index >= len(arr) || arr[index] < 0 {
		return 0
	}
	if arr[index] > 0 && arr[index]%4 == 0 {
		return arr[index] + angkalipatan(arr, index+1)
	}
	return angkalipatan(arr, index+1)
}

func main() {
	var bilangan int
	var deret []int

	fmt.Print("Masukkan serangkaian bilangan bulat dengan bilangan negatif di akhiran:")
	for {
		fmt.Scan(&bilangan)
		if bilangan < 0 {
			break
		}
		deret = append(deret, bilangan)
	}

	jumlah := angkalipatan(deret, 0)
	fmt.Println("Jumlah seluruh bilangan bulat positif (kelipatan 4) : ", jumlah)
}
