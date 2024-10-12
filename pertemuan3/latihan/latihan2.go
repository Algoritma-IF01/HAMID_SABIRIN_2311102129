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
