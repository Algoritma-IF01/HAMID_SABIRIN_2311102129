package main

import (
	"fmt"
)

func valid(voucher string) bool {
	panjang := len(voucher)
	if panjang != 5 && panjang != 6 {
		return false
	}

	duaawal := int(voucher[0]-'0') * int(voucher[1]-'0')
	duaakhir := int(voucher[panjang-2]-'0') * int(voucher[panjang-1]-'0')
	if duaawal != duaakhir {
		return false
	}

	if panjang == 5 {
		tengah := int(voucher[2] - '0')
		if tengah%2 != 0 {
			return false
		}
	} else if panjang == 6 {
		tengah1 := int(voucher[2] - '0')
		tengah2 := int(voucher[3] - '0')
		tengah3 := int(voucher[4] - '0')
		if tengah1%2 != 0 || tengah2%2 != 0 || tengah3%2 != 0 {
			return false
		}
	}

	return true
}

func main() {
	var n int
	fmt.Print("Masukan jumlah mahasiswa: ")
	fmt.Scan(&n)

	validvoucher := 0
	invalidvoucher := 0

	for i := 0; i < n; i++ {
		var voucher string
		fmt.Scan(&voucher)
		if valid(voucher) {
			validvoucher++
		} else {
			invalidvoucher++
		}
	}

	fmt.Println("Valid: ", validvoucher, "tidakvalid : ", invalidvoucher)
}
