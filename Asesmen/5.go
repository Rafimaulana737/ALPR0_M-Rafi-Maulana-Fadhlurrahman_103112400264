package main

import (
	"fmt"
)

func main() {
	// Input jumlah bola
	var n int
	fmt.Print("Masukkan jumlah bola: ")
	fmt.Scan(&n)

	// Variabel untuk menyimpan total berat
	var totalBerat float64

	// Loop untuk setiap bola
	for i := 0; i < n; i++ {
		var massaJenis, volume float64
		fmt.Printf("Masukkan massa jenis dan volume bola ke-%d: ", i+1
	}