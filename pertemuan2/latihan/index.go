package main

import "fmt"

func main() {
	var a, b, c float64
	var hipotenusa bool

	fmt.Print("Input no 1 (sisi a): ")
	fmt.Scanln(&a)

	fmt.Print("Input no 2 (sisi b): ")
	fmt.Scanln(&b)

	fmt.Print("Input no 3 (sisi c): ")
	fmt.Scanln(&c)

	hipotenusa = (c * c) == (a*a + b*b)
	fmt.Println("Sisi c adalah hipotenusa segitiga a,b,c: ", hipotenusa)
}
