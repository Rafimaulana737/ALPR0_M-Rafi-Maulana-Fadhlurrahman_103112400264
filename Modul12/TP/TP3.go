package main

import "fmt"

func main() {
	var total, harga int
	for {
		fmt.Print("Masukkan harga barang (ketik 0 untuk selesai): ")
		fmt.Scan(&harga)
		if harga == 0 {
			break
		}
		total += harga
	}
	fmt.Println("Total belanja Anda:", total)
}
