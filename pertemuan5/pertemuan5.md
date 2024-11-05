# <h1 align="center">Laporan Praktikum Modul 5 STRUCT & ARRAY</h1>

<h1 align="center">Hamid Sabirin-2311102129</h1>

<h2 align="center">PERTEMUAN 5</h2>
<h2 align="center">STRUCT & ARRAY</h2> 

### 1. Latihan1

```go
package main

import "fmt"

type bilangan int
type pecahan float64

func main() {
	var a, b bilangan
	var hasil pecahan
	a = 9
	b = 5
	hasil = pecahan(a) / pecahan(b)
	fmt.Println(hasil)

}
```

### Output Screenshot:

![latihan 1](assets/latihan1.png)

### 2. Latihan2

```go
package main

import "fmt"

type waktu struct {
    jam, menit, detik int
}

func main() {
    var wParkir, wPulang, durasi waktu
    var dParkir, dPulang, lParkir int
    
    fmt.Scan(&wParkir.jam, &wParkir.menit, &wParkir.detik)
    fmt.Scan(&wPulang.jam, &wPulang.menit, &wPulang.detik)
    
    dParkir = wParkir.detik + wParkir.menit*60 + wParkir.jam*3600
    dPulang = wPulang.detik + wPulang.menit*60 + wPulang.jam*3600
    lParkir = dPulang - dParkir
    
    durasi.jam = lParkir / 3600
    durasi.menit = lParkir % 3600 / 60
    durasi.detik = lParkir % 3600 % 60
    
	fmt.Printf("Lama Parkir: %d jam %d menit %d detik\n", durasi.jam, durasi.menit, durasi.detik)
}
```

### Output Screenshot:

![latihan 2](assets/latihan2.png)

### 3. Latihan3

```go
package main

import "fmt"

// Definisi tipe CircType
type CircType struct {
	Radius float64
}

// Definisi tipe NewType
type NewType struct {
	Name string
}

func main() {
	var (
		// Array arr mempunyai 73 elemen, masing-masing bertipe CircType
		arr [73]CircType

		// Array buf dengan 5 elemen, dengan nilai awal 7,3,5,2,11
		buf = [5]byte{7, 3, 5, 2, 11}

		// Array mhs berisi 2000 elemen NewType
		mhs [2000]NewType

		// Array dua dimensi rec berisi nilai float64
		rec [20][40]float64
	)

	// Mengisi beberapa elemen contoh
	arr[0] = CircType{Radius: 5.5}
	mhs[0] = NewType{Name: "Alice"}
	rec[0][0] = 3.14

	// Contoh penggunaan variabel
	fmt.Println("arr[0]:", arr[0])
	fmt.Println("buf:", buf)
	fmt.Println("mhs[0]:", mhs[0])
	fmt.Println("rec[0][0]:", rec[0][0])
}
```

### Output Screenshot:

![latihan 3](assets/latihan3.png)

### 4. Latihan4

```go
package main

import "fmt"

func main() {
	// Membuat map dengan tipe string sebagai kunci dan integer sebagai nilai
	scores := map[string]int{
		"John": 90,
		"Anne": 85,
	}

	// Mengambil nilai dari kunci
	johnScore := scores["John"]
	fmt.Println("Nilai John:", johnScore)

	// Mengganti nilai dari kunci yang ada
	scores["John"] = 95
	fmt.Println("Nilai John setelah diganti:", scores["John"])

	// Menambah kunci baru dengan nilai
	scores["David"] = 88
	fmt.Println("Nilai David ditambahkan:", scores["David"])

	// Menghapus kunci dari map
	delete(scores, "Anne")
	fmt.Println("Map setelah kunci 'Anne' dihapus:", scores)

	// Mengecek apakah kunci ada dalam map
	if score, ada := scores["David"]; ada {
		fmt.Println("Nilai David ditemukan:", score)
	} else {
		fmt.Println("Nilai David tidak ditemukan")
	}
}
```

### Output Screenshot:

![latihan 4](assets/latihan4.png)

### 5. Latihan5

```go
package main

import (
	"fmt"
	"math"
)

// Definisi tipe bentukan untuk titik
type Titik struct {
	x int
	y int
}

// Definisi tipe bentukan untuk lingkaran
type Lingkaran struct {
	center Titik
	radius int
}

// Fungsi untuk menghitung jarak antara dua titik
func jarak(p Titik, q Titik) float64 {
	return math.Sqrt(float64((p.x-q.x)*(p.x-q.x) + (p.y-q.y)*(p.y-q.y)))
}

// Fungsi untuk menentukan apakah titik berada di dalam lingkaran
func didalam(c Lingkaran, p Titik) bool {
	return jarak(p, c.center) < float64(c.radius)
}

func main() {
	var (
		// Mengambil input untuk lingkaran 1
		lingkaran1 Lingkaran
		// Mengambil input untuk lingkaran 2
		lingkaran2 Lingkaran
		// Mengambil input untuk titik sembarang
		point Titik
	)

	// Input untuk lingkaran 1 (cx, cy, r)
	fmt.Println("Masukkan koordinat titik pusat dan radius lingkaran 1 (cx cy r):")
	fmt.Scan(&lingkaran1.center.x, &lingkaran1.center.y, &lingkaran1.radius)

	// Input untuk lingkaran 2 (cx, cy, r)
	fmt.Println("Masukkan koordinat titik pusat dan radius lingkaran 2 (cx cy r):")
	fmt.Scan(&lingkaran2.center.x, &lingkaran2.center.y, &lingkaran2.radius)

	// Input untuk titik sembarang (x, y)
	fmt.Println("Masukkan koordinat titik sembarang (x y):")
	fmt.Scan(&point.x, &point.y)

	// Mengecek posisi titik terhadap kedua lingkaran
	inLingkaran1 := didalam(lingkaran1, point)
	inLingkaran2 := didalam(lingkaran2, point)

	if inLingkaran1 && inLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if inLingkaran1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if inLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
```

### Output Screenshot:

![latihan 5](assets/latihan5.png)

### 6. Latihan6

```go
package main  

import (  
	"fmt"  
	"math"  
)  

func main() {  

	type titik struct {  
		x, y float64  
	}  

	var a, b, c, d titik  

	fmt.Scanf("%f %f\n", &a.x, &a.y)  
	fmt.Scanf("%f %f\n", &b.x, &b.y)  
	fmt.Scanf("%f %f\n", &c.x, &c.y)  
	fmt.Scanf("%f %f\n", &d.x, &d.y)  

	panjang := math.Abs(a.x - b.x)  
	lebar := math.Abs(a.y - d.y)  

	kelilingPersegiPanjang := 2 * (panjang + lebar)  

	fmt.Print(kelilingPersegiPanjang)  
}
```

### Output Screenshot:

![latihan 6](assets/latihan6.png)

### 7. Tugas1

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah elemen array (N): ")
	fmt.Scanln(&n)

	array := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan elemen ke-%d: ", i)
		fmt.Scanln(&array[i])
	}

	var choice, x, index int
	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Menampilkan keseluruhan isi dari array.")
		fmt.Println("2. Menampilkan elemen-elemen array dengan indeks ganjil saja.")
		fmt.Println("3. Menampilkan elemen-elemen array dengan indeks genap saja.")
		fmt.Println("4. Menampilkan elemen-elemen array dengan indeks kelipatan bilangan x.")
		fmt.Println("5. Menghapus elemen array pada indeks tertentu.")
		fmt.Println("6. Menampilkan rata-rata dari bilangan yang ada di dalam array.")
		fmt.Println("7. Menampilkan standar deviasi dari bilangan yang ada di dalam array.")
		fmt.Println("8. Menampilkan frekuensi dari setiap bilangan tertentu di dalam array.")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih opsi: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("Isi array:", array)

		case 2:
			fmt.Print("Elemen dengan indeks ganjil: ")
			for i := 1; i < len(array); i += 2 {
				fmt.Print(array[i], " ")
			}
			fmt.Println()

		case 3:
			fmt.Print("Elemen dengan indeks genap: ")
			for i := 0; i < len(array); i += 2 {
				fmt.Print(array[i], " ")
			}
			fmt.Println()

		case 4:
			fmt.Print("Masukkan nilai x: ")
			fmt.Scanln(&x)
			fmt.Printf("Elemen dengan indeks kelipatan %d: ", x)
			for i := x; i < len(array); i += x {
				fmt.Print(array[i], " ")
			}
			fmt.Println()

		case 5:
			fmt.Print("Masukkan indeks yang ingin dihapus: ")
			fmt.Scanln(&index)
			if index >= 0 && index < len(array) {
				array = append(array[:index], array[index+1:]...)
				fmt.Println("Isi array setelah penghapusan:", array)
			} else {
				fmt.Println("Indeks tidak valid!")
			}

		case 6:
			total := 0
			for _, value := range array {
				total += value
			}
			avg := float64(total) / float64(len(array))
			fmt.Printf("Rata-rata dari array: %.2f\n", avg)

		case 7:
			total := 0
			for _, value := range array {
				total += value
			}
			mean := float64(total) / float64(len(array))
			variance := 0.0
			for _, value := range array {
				variance += math.Pow(float64(value)-mean, 2)
			}
			stdDev := math.Sqrt(variance / float64(len(array)))
			fmt.Printf("Standar deviasi dari array: %.2f\n", stdDev)

		case 8:
			frequency := make(map[int]int)
			for _, value := range array {
				frequency[value]++
			}
			fmt.Println("Frekuensi setiap bilangan dalam array:")
			for num, freq := range frequency {
				fmt.Printf("%d: %d kali\n", num, freq)
			}

		case 0:
			fmt.Println("Terima kasih!")
			return

		default:
			fmt.Println("Pilihan tidak valid, silakan coba lagi.")
		}
	}
}
```

### Output Screenshot:

![tugas 1](assets/soal1a.png)
![tugas 1](assets/soal1b.png)
![tugas 1](assets/soal1c.png)
![tugas 1](assets/soal1d.png)
![tugas 1](assets/soal1e.png)
![tugas 1](assets/soal1f.png)
![tugas 1](assets/soal1g.png)
![tugas 1](assets/soal1h.png)

### 8. Tugas2

```go
package main

import (
	"fmt"
)

func main() {
	var clubA, clubB string
	var scoreA, scoreB int
	var numMatches int
	var winners []string

	fmt.Print("Masukkan nama Klub A: ")
	fmt.Scanln(&clubA)
	fmt.Print("Masukkan nama Klub B: ")
	fmt.Scanln(&clubB)
	fmt.Print("Masukkan jumlah pertandingan: ")
	fmt.Scanln(&numMatches)

	for i := 1; i <= numMatches; i++ {
		fmt.Printf("Pertandingan %d (skor %s): ", i, clubA)
		fmt.Scanln(&scoreA)
		fmt.Printf("Pertandingan %d (skor %s): ", i, clubB)
		fmt.Scanln(&scoreB)

		if scoreA < 0 || scoreB < 0 {
			fmt.Println("Pertandingan selesai")
			break
		}

		if scoreA > scoreB {
			winners = append(winners, clubA)
		} else if scoreB > scoreA {
			winners = append(winners, clubB)
		} else {
			winners = append(winners, "Draw")
		}
	}

	fmt.Println("\nHasil Pertandingan:")
	for i, winner := range winners {
		if winner == "Draw" {
			fmt.Printf("Hasil %d : Draw\n", i+1)
		} else {
			fmt.Printf("Hasil %d : %s\n", i+1, winner)
		}
	}
	fmt.Println("Pertandingan selesai")
}
```

### Output Screenshot:

![tugas 2](assets/soal2.png)

### 9. Tugas3 

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(tab *tabel, n *int) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan karakter dalam satu baris (akhiri dengan titik '.') : ")
	input, _ := reader.ReadString('\n')

	for i, char := range input {
		if char == '.' || i >= NMAX {
			break
		}
		if char != ' ' && char != '\n' {
			tab[*n] = char
			*n++
		}
	}
}

func cetakArray(tab tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c ", tab[i])
	}
	fmt.Println()
}

func balikkanArray(tab *tabel, n int) {
	for i := 0; i < n/2; i++ {
		tab[i], tab[n-i-1] = tab[n-i-1], tab[i]
	}
}

func palindrom(tab tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if tab[i] != tab[n-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var n int

	isiArray(&tab, &n)

	// Cetak teks asli
	fmt.Print("Teks: ")
	cetakArray(tab, n)

	// Cek apakah teks asli adalah palindrom
	if palindrom(tab, n) {
		fmt.Println("Palindrom: true")
	} else {
		fmt.Println("Palindrom: false")
	}

	// Membalikkan isi array dan cetak teks setelah dibalik
	balikkanArray(&tab, n)
	fmt.Print("Reverse teks: ")
	cetakArray(tab, n)
}
```

### Output Screenshot:

![tugas 3](assets/soal3.png)