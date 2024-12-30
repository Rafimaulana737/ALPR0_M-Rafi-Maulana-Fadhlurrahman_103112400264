Mulai
    Tanyakan "Masukkan jumlah bilangan: "
    Baca n  // jumlah bilangan yang akan dimasukkan

    Tanyakan "Masukkan bilangan pertama: "
    Baca bilangan_terbesar
    untuk i ← 2 hingga n:
        Tanyakan "Masukkan bilangan berikutnya: "
        Baca bilangan

        jika bilangan > bilangan_terbesar:
            bilangan_terbesar ← bilangan  // Update bilangan terbesar

    Tampilkan "Bilangan terbesar adalah: " bilangan_terbesar
Selesai

package main

import "fmt"

func main() {
	var bilangan int
	var digitTerbesar int

	// Input bilangan bulat positif
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&bilangan)

	// Proses untuk mencari digit terbesar
	for bilangan > 0 {
		digit := bilangan % 10
		if digit > digitTerbesar {
			digitTerbesar = digit
		}
		bilangan = bilangan / 10
	}

	// Output digit terbesar
	fmt.Printf("Digit terbesar: %d\n", digitTerbesar)
}
