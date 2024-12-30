package main

import "fmt"

func main() {
	var S, minggu int

	// Input jatah berenang
	fmt.Print("Masukkan jatah berenang maksimum: ")
	fmt.Scan(&S)

	// Menghitung minggu ke berapa jatah berenang habis
	minggu = S / 7
	if S%7 != 0 {
		minggu += 1
	}

	// Output minggu ke berapa
	fmt.Printf("Minggu ke-%d\n", minggu)
}
