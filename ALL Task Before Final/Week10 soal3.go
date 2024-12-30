Mulai
    Tanyakan "Masukkan keuntungan bulan ini: "
    Baca keuntungan_bulan_ini

    Tanyakan "Masukkan keuntungan bulan sebelumnya: "
    Baca keuntungan_bulan_sebelumnya

    keuntungan_berbeda ← keuntungan_bulan_ini - keuntungan_bulan_sebelumnya

    jika keuntungan_berbeda > 0:
        Tampilkan "Peningkatan keuntungan sebesar" keuntungan_berbeda
    jika keuntungan_berbeda < 0:
        Tampilkan "Penurunan keuntungan sebesar" (-keuntungan_berbeda)
    jika keuntungan_berbeda = 0:
        Tampilkan "Keuntungan tetap"
Selesai



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
