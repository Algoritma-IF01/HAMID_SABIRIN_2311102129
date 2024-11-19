// package main

// import (
// 	"fmt"
// )

// type arrInt [2023]int

// func terkecil_1(tabInt arrInt, n int) int {
// 	var min int = tabInt[0]
// 	var idx int = 1

// 	for j := 1; j < n; j++ {
// 		if tabInt[idx] < min {
// 			min = tabInt[idx]
// 		}
// 		idx = j + 1
// 	}
// 	return min
// }

// func main() {
// 	var n int
// 	fmt.Print("Masukkan jumlah elemen: ")
// 	fmt.Scan(&n)

// 	var data arrInt
// 	fmt.Println("Masukkan elemen-elemen array:")
// 	for i := 0; i < n; i++ {
// 		fmt.Scan(&data[i])
// 	}

// 	smallest := terkecil_1(data, n)
// 	fmt.Printf("Nilai terkecil adalah: %d\n", smallest)
// }

package main

import (
	"fmt"
)

type arrInt [2023]int

// Fungsi untuk mencari indeks dari nilai terkecil
func terkecil_2(tabInt arrInt, n int) int {
	var idx int = 0
	var j int = 1
	for j < n {
		if tabInt[idx] > tabInt[j] {
			idx = j
		}
		j = j + 1
	}
	return idx
}

func main() {
	var n int
	var data arrInt

	// Input jumlah elemen N
	fmt.Print("Masukkan jumlah elemen (N <= 2023): ")
	fmt.Scan(&n)

	// Validasi N agar tidak melebihi kapasitas array
	if n <= 0 || n > 2023 {
		fmt.Println("Jumlah elemen harus antara 1 dan 2023")
		return
	}

	// Input elemen-elemen array
	fmt.Println("Masukkan elemen array:")
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}

	// Panggil fungsi untuk mencari indeks nilai terkecil
	idxTerkecil := terkecil_2(data, n)
	fmt.Printf("Indeks nilai terkecil: %d\n", idxTerkecil)
	fmt.Printf("Nilai terkecil: %d\n", data[idxTerkecil])
}
