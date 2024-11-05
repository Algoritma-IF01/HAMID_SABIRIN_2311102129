package main

func main() {
	for i := 0; i < 10; i++ {
		if i == 4 {
			continue
		}
		println(i)
	}
}
