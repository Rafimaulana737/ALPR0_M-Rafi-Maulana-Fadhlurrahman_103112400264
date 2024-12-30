package main

import "fmt"

func main() {
	var keuntunganBulanIni, keuntunganBulanSebelumnya float64

	// Input keuntungan bulan ini dan bulan sebelumnya
	fmt.Print("Masukkan keuntungan bulan ini: ")
	fmt.Scan(&keuntunganBulanIni)
	fmt.Print("Masukkan keuntungan bulan sebelumnya: ")
	fmt.Scan(&keuntunganBulanSebelumnya)

	// Menghitung selisih keuntungan
	kenaikan := keuntunganBulanIni - keuntunganBulanSebelumnya

	// Menentukan apakah ada peningkatan atau penurunan keuntungan
	if kenaikan > 0 {
		fmt.Printf("Peningkatan keuntungan sebesar %.2f\n", kenaikan)
	} else if kenaikan < 0 {
		fmt.Printf("Penurunan keuntungan sebesar %.2f\n", -kenaikan)
	} else {
		fmt.Println("Keuntungan tetap")
	}
}
