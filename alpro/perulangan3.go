package main

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			println("Nilainya sudah 5, berhenti")
			break
		}
		println(i)
	}
}
