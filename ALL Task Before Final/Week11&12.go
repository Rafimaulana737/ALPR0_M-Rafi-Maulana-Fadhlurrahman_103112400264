Mulai
    saldo ← 0  // Inisialisasi saldo awal

    ulangi selama transaksi tidak 0:
        Tanyakan "Masukkan transaksi: "
        Baca transaksi

        jika transaksi > 0:
            saldo ← saldo + transaksi  // Uang masuk
        jika transaksi < 0:
            saldo ← saldo - transaksi  // Uang keluar

    Tampilkan saldo
Selesai







package main

import "fmt"

func main() {
	var transaksi, saldo int

	// Inisialisasi saldo awal
	saldo = 0
	
	// Input transaksi berulang
	for {
		// Masukkan transaksi
		fmt.Print("Masukkan transaksi: ")
		fmt.Scan(&transaksi)

		// Jika transaksi adalah 0, keluar dari perulangan
		if transaksi == 0 {
			break
		}

		// Tambahkan atau kurangi saldo berdasarkan transaksi
		saldo += transaksi
	}

	// Output saldo akhir
	fmt.Printf("Saldo akhir: %d\n", saldo)
}
