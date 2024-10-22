package main

import "fmt"

//fungsi rekursif untuk deret fibonancci
func fibonanci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibonanci(n-1) + fibonanci(n-2)
	}
}

func main() {
	// output deret fibonanci hingga suku ke 1
	fmt.Println("Deret fibonanci hingga suku ke 10 :")
	for i := 0; i <= 10; i++ {
		fmt.Printf("fibonanci(%d)=%d\n", i, fibonanci(i))
	}
}
