package main

import (
	"fmt"
)

func main() {
	var jumlahIkan, kapasitasWadah int
	fmt.Print("Masukkan jumlah ikan (x) dan kapasitas per wadah (y): ")
	fmt.Scan(&jumlahIkan, &kapasitasWadah)

	// Input berat ikan
	beratIkan := make([]float64, jumlahIkan)
	fmt.Println("Masukkan berat ikan (dipisah dengan spasi):")
	for i := 0; i < jumlahIkan; i++ {
		fmt.Scan(&beratIkan[i])
	}

	// Membagi ikan ke dalam wadah
	wadah := [][]float64{}
	wadahSaatIni := []float64{}
	jumlahDalamWadah := 0

	for i := 0; i < jumlahIkan; i++ {
		if jumlahDalamWadah < kapasitasWadah {
			wadahSaatIni = append(wadahSaatIni, beratIkan[i])
			jumlahDalamWadah++
		} else {
			wadah = append(wadah, wadahSaatIni)
			wadahSaatIni = []float64{beratIkan[i]}
			jumlahDalamWadah = 1
		}
	}
	// Menambahkan wadah terakhir
	if len(wadahSaatIni) > 0 {
		wadah = append(wadah, wadahSaatIni)
	}

	// Menghitung total berat dan rata-rata berat setiap wadah
	totalBeratPerWadah := []float64{}
	for _, satuWadah := range wadah {
		jumlahBerat := 0.0
		for _, berat := range satuWadah {
			jumlahBerat += berat
		}
		totalBeratPerWadah = append(totalBeratPerWadah, jumlahBerat)
	}

	// Output hasil
	fmt.Println("Output:")
	// Baris pertama: Total berat ikan di setiap wadah
	for _, total := range totalBeratPerWadah {
		fmt.Printf("%.2f ", total)
	}
	fmt.Println()

	// Baris kedua: Rata-rata berat ikan di setiap wadah
	rataRata := 0.0
	for _, total := range totalBeratPerWadah {
		rataRata += total
	}
	rataRata /= float64(len(totalBeratPerWadah))
	fmt.Printf("%.2f\n", rataRata)
}
