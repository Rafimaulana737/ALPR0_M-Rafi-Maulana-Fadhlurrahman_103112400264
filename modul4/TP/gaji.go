package main

import (
	"fmt"
)

func main() {
	// Input dari pengguna
	var jamKerja float64
	var upahPerJam float64

	fmt.Print("Masukkan jumlah jam kerja dalam seminggu: ")
	fmt.Scan(&jamKerja)
	fmt.Print("Masukkan upah per jam: ")
	fmt.Scan(&upahPerJam)

	// Menghitung total gaji mingguan
	var totalGaji float64
	if jamKerja > 40 {
		lembur := jamKerja - 40
		totalGaji = (40 * upahPerJam) + (lembur * upahPerJam * 1.5)
	} else {
		totalGaji = jamKerja * upahPerJam
	}

	// Menghitung total gaji bulanan (4 minggu)
	totalGajiBulanan := totalGaji * 4

	// Menampilkan hasil
	fmt.Printf("Total gaji bulanan: %.2f\n", totalGajiBulanan)
}
