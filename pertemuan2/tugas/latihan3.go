package main

import (
	"fmt"
	"math"
)

func main() {
	var radius float64
	fmt.Print("Jejari = ")
	fmt.Scan(&radius)

	const phi = 3.1415926535
	volume := (4.0 / 3.0) * phi * math.Pow(radius, 3)
	surfaceArea := 4 * phi * math.Pow(radius, 2)

	fmt.Printf("Bola dengan jejari %.0f memiliki volume %.4f dan luas kulit %.4f\n", radius, volume, surfaceArea)
}
