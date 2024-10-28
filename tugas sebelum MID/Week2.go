package main

import "fmt"

func main() {
	var radius float64
	fmt.Print("Masukkan jari-jari lingkaran: ")
	fmt.Scan(&radius)

	luas := 3.14 * radius * radius
	keliling := 2 * 3.14 * radius

	fmt.Printf("Luas lingkaran: %.2f\n", luas)
	fmt.Printf("Keliling lingkaran: %.2f\n", keliling)
}
