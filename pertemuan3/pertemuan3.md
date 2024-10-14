# <h1 align="center">Laporan Praktikum Modul 3 PERULANGAN & PERCABANGAN</h1>

<h1 align="center">Hamid Sabirin-2311102129</h1>

<h2 align="center">PERTEMUAN 3</h2>
<h2 align="center">PERULANGAN & PERCABANGAN</h2>

### 1. Latihan1

```go
package main

import "fmt"

func main() {
	urutanBenar := []string{"merah", "kuning", "hijau", "ungu"}
	hasil := true

	for i := 1; i <= 5; i++ {
		var warna1, warna2, warna3, warna4 string
		fmt.Printf("percobaan %d\n ", i)
		fmt.Print("masukkan warna pertama: ")
		fmt.Scan(&warna1)
		fmt.Print("masukkan warna kedua: ")
		fmt.Scan(&warna2)
		fmt.Print("masukkan warna ketiga: ")
		fmt.Scan(&warna3)
		fmt.Print("masukkan warna keempat: ")
		fmt.Scan(&warna4)

		if warna1 != urutanBenar[0] || warna2 != urutanBenar[1] || warna3 != urutanBenar[2] || warna4 != urutanBenar[3] {
			hasil = false
		}
	}
	println("BERHASIL", hasil)
}
```

### Output Screenshot:

![latihan 1](assets/latihan1.png)

### 2. Latihan2

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var pita string
	var bungaCount int

	for {
		fmt.Printf("bunga %d: ", bungaCount+1)
		scanner.Scan()
		input := scanner.Text()
		if strings.ToLower(input) == "selesai" {
			break
		}

		if pita == "" {
			pita = input
		} else {
			pita += "-" + input
		}
		bungaCount++
	}
	fmt.Printf("pita : %s\n", pita)
	fmt.Printf("bunga : %d\n", bungaCount)
}
```

### Output Screenshot:

![latihan 2](assets/latihan2.png)

### 3. Tugas1

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	var kantong1, kantong2 float64
	oleng := false

	for {
		fmt.Print("Masukkan berat belanjaan di kedua kantong: ")
		fmt.Scan(&kantong1, &kantong2)

		if kantong1+kantong2 > 150 {
			fmt.Println("Berat melebihi 150")
			fmt.Println("Proses selesai.")
			break
		}

		if kantong1 < 0 || kantong2 < 0 {
			fmt.Println("Proses selesai.")
			break
		}

		selisih := math.Abs(kantong1 - kantong2)
		oleng = selisih > 9

		fmt.Printf("Sepeda motor Pak Andi akan oleng: %v\n", oleng)
	}
}
```

### Output Screenshot:

![tugas 1](assets/tugas1.png)

### 4. Tugas2

```go
package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menghitung f(k)
func f(k float64) float64 {
	numerator := math.Pow((4*k + 2), 2)
	denominator := (4*k + 1) * (4*k + 3)
	return numerator / denominator
}

func sqrt2(k int) float64 {
	result := 1.0
	for i := 0; i <= k; i++ {
		result *= f(float64(i))
	}
	return result
}

func main() {
	var K int

	for i := 1; i <= 3; i++ {
		fmt.Print("Nilai K = ")
		fmt.Scan(&K)

		approxSqrt2 := sqrt2(K)
		fmt.Printf("Nilai akar 2 = %.10f\n\n", approxSqrt2)
	}

	fmt.Println("Proses selesai.")
}
```

### Output Screenshot:

![tugas 2](assets/tugas2.png)

### 5. Tugas3

```go
package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 3; i++ {
		var beratParsel, kg, gram int
		var biayaPerKg, biayaSisa, totalBiaya int

		// Input berat parsel dalam gram
		fmt.Printf("Contoh #%d:\n ", i)
		fmt.Print("Berat parsel (gram): ")
		fmt.Scan(&beratParsel)

		// berat dalam kg dan sisa gram
		kg = beratParsel / 1000
		gram = beratParsel % 1000

		//  biaya per kg 
		biayaPerKg = kg * 10000

		// biaya sisa berdasarkan gram
		if gram > 0 && kg < 10 {
			if gram <= 500 {
				biayaSisa = gram * 5
			} else {
				biayaSisa = gram * 15
			}
		} else {
			biayaSisa = 0 
		}

		// total biaya
		totalBiaya = biayaPerKg + biayaSisa

		// Output 
		fmt.Printf("Detail berat: %d kg + %d gram\n", kg, gram)
		fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaPerKg, biayaSisa)
		fmt.Printf("Total biaya: Rp. %d\n\n", totalBiaya)
	}

	fmt.Println("Proses selesai.")
}
```

### Output Screenshot:

![tugas 3](assets/tugas3.png)

### 6. Tugas4a

```go
// Sebelum Di Perbaiki
package main

import "fmt"

func main() {
    var nam float64
    var nmk string
    fmt.Print("Nilai akhir mata kuliah: ")
    fmt.Scanln(&nam)

    if nam > 80 {
        nam = "A"
    }
    if nam > 72.5 {
        nam = "AB"
    }
    if nam > 65 {
        nam = "B"
    }
    if nam > 57.5 {
        nam = "BC"
    }
    if nam > 50 {
        nam = "C"
    }
    if nam > 40 {
        nam = "D"
    } else if nam <= 40 {
        nam = "E"
    }

    fmt.Println("Nilai mata kuliah: ", nmk)
}

```

### Output Screenshot:

![tugas 4a](assets/tugas4a.png)

### 6. Tugas4b

A. Daftar Kesalahan:

1. Penggunaan variabel yang salah (nam diubah menjadi string):
- Pada baris seperti nam = "A", "AB", "B", dll., program mencoba menyimpan nilai string ke dalam variabel nam yang sebelumnya dideklarasikan sebagai float64. Variabel nam seharusnya digunakan untuk nilai numerik, bukan string. Variabel yang seharusnya digunakan untuk menyimpan nilai huruf (nilai akhir mata kuliah) adalah nmk.
Perbaikan: Gantilah nam = "A" menjadi nmk = "A".

2. Output menggunakan variabel yang salah (nmk tidak pernah diisi):
- Program mencetak variabel nmk di akhir, tetapi variabel ini tidak pernah diisi dengan nilai apapun. Nilai-nilai huruf seperti "A", "B", dan seterusnya seharusnya disimpan dalam nmk, bukan dalam nam.
Perbaikan: Pastikan semua penugasan nilai string (seperti "A", "B", dll.) disimpan ke dalam variabel nmk, bukan nam.

3. Kondisi terpisah tanpa else if:
- Semua pernyataan if dievaluasi secara terpisah. Artinya, meskipun nam > 80 benar, semua kondisi if berikutnya tetap diperiksa, sehingga nilai yang benar bisa tertimpa oleh kondisi berikutnya.
Perbaikan: Gunakan struktur else if untuk memastikan bahwa hanya satu kondisi yang dieksekusi berdasarkan rentang nilai nam.

4. Masalah logika pada kondisi if nam <= 40:
- Pada blok terakhir, ada kombinasi if nam > 40 dan else if nam <= 40. Ini secara teknis benar, tetapi bisa lebih sederhana. Jika semua kondisi sebelumnya salah (yaitu, semua nilai nam lebih kecil atau sama dengan 40), cukup gunakan else untuk menangani kondisi tersebut.
Perbaikan: Gantilah else if nam <= 40 dengan else.

B. Mengapa demikian ? 

1. Tipe Data Tidak Sesuai

Kesalahan: 
Program mencoba menyimpan nilai string (seperti "A", "AB", dll.) ke dalam variabel nam yang bertipe float64. Ini tidak diperbolehkan dalam Go karena tipe data harus konsisten. float64 hanya dapat menyimpan nilai numerik, bukan string.

Mengapa demikian: 
Go adalah bahasa yang statically typed, yang berarti variabel tidak dapat berubah tipe datanya setelah dideklarasikan. Oleh karena itu, menyimpan string dalam variabel bertipe numerik (float64) akan menimbulkan error.

2. Urutan Logika yang Tidak Efisien (Tanpa else if)
Kesalahan: 
Menggunakan if secara berurutan tanpa else if menyebabkan semua kondisi dievaluasi, meskipun sudah ada satu kondisi yang benar. Ini bisa membuat program mengevaluasi kondisi yang tidak perlu dan menyebabkan nilai nmk berubah tidak sesuai.

Mengapa demikian: 
Ketika hanya menggunakan if, setiap kondisi akan diperiksa, meskipun sebelumnya sudah ada kondisi yang terpenuhi. Ini membuat logika program menjadi tidak efisien, karena seharusnya cukup memeriksa satu kondisi yang benar dan langsung berhenti.

3. Variabel nmk Tidak Pernah Diinisialisasi
Kesalahan: 
Variabel nmk dideklarasikan tetapi tidak pernah diisi atau diinisialisasi dengan nilai. Akibatnya, ketika program mencoba mencetak nilai nmk, hasilnya adalah string kosong.

Mengapa demikian: 
Setiap variabel dalam Go harus diinisialisasi atau diisi nilainya sebelum digunakan. Jika tidak, hasilnya bisa tidak sesuai dengan yang diharapkan (seperti mencetak string kosong).

C. Alur Program Seharusnya :

Program seharusnya menanyakan nilai akhir mata kuliah dari pengguna dan kemudian menentukan nilai huruf berdasarkan rentang nilai tersebut. 
Berikut adalah alur program yang benar:
1. Menerima input nilai akhir mata kuliah.
2. Menggunakan else if untuk memastikan hanya satu kondisi yang dievaluasi.
3. Menggunakan variabel terpisah untuk menyimpan nilai huruf (nmk).
4. Mencetak nilai huruf berdasarkan nilai numerik.

### 7. Tugas4c

```go
// Sesudah Di Perbaiki
package main

import "fmt"

func main() {
	var nam float64
	var nmk string

	fmt.Print("Nilai akhir mata kuliah: ")
	fmt.Scanln(&nam)

	if nam > 80 {
		nmk = "A"
	} else if nam > 72.5 {
		nmk = "AB"
	} else if nam > 65 {
		nmk = "B"
	} else if nam > 57.5 {
		nmk = "BC"
	} else if nam > 50 {
		nmk = "C"
	} else if nam > 40 {
		nmk = "D"
	} else {
		nmk = "E"
	}

	fmt.Println("Nilai mata kuliah: ", nmk)
}
```

### Output Screenshot:

![tugas 4c](assets/tugas4c.png)

### 8. Tugas5

```go
package main

import "fmt"

func main() {
	for i := 1; i <= 2; i++ {
		var b int

		fmt.Printf("Bilangan ke-%d: ", i)
		fmt.Scan(&b)

		fmt.Print("Faktor: ")
		var jumlahFaktor int
		for j := 1; j <= b; j++ {
			if b%j == 0 {
				fmt.Printf("%d ", j)
				jumlahFaktor++
			}
		}
		fmt.Println()

		isPrima := jumlahFaktor == 2

		fmt.Printf("Prima: %t\n\n", isPrima)
	}

	fmt.Println("Proses selesai.")
}
```

### Output Screenshot:

![tugas 5](assets/tugas5.png)
