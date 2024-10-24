package main

import "fmt"

func main() {
    var jumlahBarang int
    fmt.Print("Masukkan jumlah barang yang dibeli: ")
    fmt.Scan(&jumlahBarang)

    poin := jumlahBarang * 10
    if jumlahBarang > 5 {
        poin += (jumlahBarang - 5) * 5
    }

    fmt.Printf("Total poin yang didapatkan: %d poin\n", poin)
}
