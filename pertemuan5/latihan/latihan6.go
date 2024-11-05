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
