package main

import (
	"fmt"
)

func main() {
	// Input nilai a dan b
	var a, b int
	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)

	// Variabel untuk menyimpan hasil penjumlahan
	sum := 0

	// Melakukan iterasi dari a sampai b
	for i := a; i <= b; i++ {
		// Mengecek apakah bilangan ganjil
		if i%2 != 0 {
			sum += i
		}
	}

	// Output hasil penjumlahan
	fmt.Printf("Hasil penjumlahan bilangan ganjil antara %d dan %d adalah: %d\n", a, b, sum)
}