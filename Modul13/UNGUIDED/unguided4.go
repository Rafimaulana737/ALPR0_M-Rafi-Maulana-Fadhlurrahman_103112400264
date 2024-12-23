package main

import "fmt"

func main() {
	var namabunga string
	var pita string
	var jumlahBunga int

	jumlahBunga = 0

	for {
		fmt.Print("Bunga", jumlahBunga+1, ": ")
		fmt.Scan(&namabunga)

		if namabunga == "SELESAI" {
			break
		}

		pita = pita + namabunga + " - "
		jumlahBunga++
	}
	fmt.Println("Pita: ", pita)
	fmt.Println("Bunga: ", jumlahBunga)
}
