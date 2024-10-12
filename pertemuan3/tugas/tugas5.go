package main

import "fmt"

func main() {
	for i := 1; i <= 2; i++ {
		var b int

		// Input bilangan bulat
		fmt.Printf("Bilangan ke-%d: ", i)
		fmt.Scan(&b)

		// menampilkan faktor-faktor dari b
		fmt.Print("Faktor: ")
		var jumlahFaktor int
		for j := 1; j <= b; j++ {
			if b%j == 0 {
				fmt.Printf("%d ", j)
				jumlahFaktor++
			}
		}
		fmt.Println()

		// apakah bialnagn prima?
		isPrima := jumlahFaktor == 2

		// print apakah bilangan prima atau bukan
		fmt.Printf("Prima: %t\n\n", isPrima)
	}

	fmt.Println("Proses selesai.")
}