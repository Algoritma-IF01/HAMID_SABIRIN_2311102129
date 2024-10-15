package main

import (
	"fmt"
)

func hitungBiayaPos(berat int) (int, int, int, int, bool) {
	biayaPerKg := 10000
	kg := berat / 1000
	sisaGram := berat % 1000
	biayaKg := kg * biayaPerKg
	var biayaSisaGram int
	biayaDitambahkan := true 

	// jika berat total >= 10kg, sisa gram tetap dihitung, tapi tidak ditambahkan ke total biaya
	if kg >= 10 {
		biayaSisaGram = sisaGram * 5 
		biayaDitambahkan = false    
	} else {
		// menghitung biaya untuk sisa gram jika berat kurang dari 10kg
		if sisaGram >= 500 {
			biayaSisaGram = sisaGram * 5 
		} else {
			biayaSisaGram = sisaGram * 15 
		}
	}

	return kg, sisaGram, biayaKg, biayaSisaGram, biayaDitambahkan
}

func main() {
	for i := 1; i <= 3; i++ {
		var berat int
		fmt.Printf("Contoh #%d:\n", i)
		fmt.Print("Berat parsel (gram): ")
		fmt.Scan(&berat)

		kg, sisaGram, biayaKg, biayaSisaGram, biayaDitambahkan := hitungBiayaPos(berat)
		totalBiaya := biayaKg
		if biayaDitambahkan {
			totalBiaya += biayaSisaGram
		}

		fmt.Printf("Detail berat: %d kg + %d gr\n", kg, sisaGram)
		fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaKg, biayaSisaGram)
		fmt.Printf("Total biaya: Rp. %d\n\n", totalBiaya)
	}
}
