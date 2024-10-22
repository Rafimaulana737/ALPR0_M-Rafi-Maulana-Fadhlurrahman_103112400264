package main

import (
	"fmt"
	"strings"
)

func main() {
	// Input jumlah nilai boolean
	var n int
	fmt.Print("Masukkan jumlah nilai boolean: ")
	fmt.Scan(&n)

	// Input nilai boolean dalam bentuk string
	var input string
	fmt.Print("Masukkan nilai boolean (true/false dipisahkan spasi): ")
	fmt.Scanln(&input)
	fmt.Scanln(&input) // Untuk menangkap masukan lebih dari satu kata

	// Memisahkan nilai boolean
	boolValues := strings.Split(input, " ")

	// Variabel untuk menyimpan hasil operasi AND
	result := true

	// Melakukan operasi AND untuk semua nilai boolean
	for i := 0; i < n; i++ {
		if boolValues[i] == "false" {
			result = false
			break
		}
	}

	// Output hasil
	fmt.Printf("Hasil operasi AND: %v\n", result)
}