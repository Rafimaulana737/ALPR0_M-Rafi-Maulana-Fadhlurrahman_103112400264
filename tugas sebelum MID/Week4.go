package main

import "fmt"

func main() {
	var nama string
	var gajiPokok, tunjangan, potongan float64

	fmt.Print("Masukkan nama: ")
	fmt.Scanln(&nama)
	fmt.Print("Masukkan gaji pokok: ")
	fmt.Scanln(&gajiPokok)
	fmt.Print("Masukkan tunjangan: ")
	fmt.Scanln(&tunjangan)
	fmt.Print("Masukkan potongan: ")
	fmt.Scanln(&potongan)

	totalGaji := gajiPokok + tunjangan - potongan
	fmt.Printf("Total gaji karyawan: %.2f\n", totalGaji)
}
