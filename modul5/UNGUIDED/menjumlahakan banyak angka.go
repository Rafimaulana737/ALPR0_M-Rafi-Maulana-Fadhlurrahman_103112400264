package main

import "fmt"

func main() {
    var n, total int
    fmt.Print("Masukkan bilangan bulat positif n: ")
    fmt.Scan(&n)

    for i := 1; i <= n; i++ {
        total += i
    }

    fmt.Printf("Hasil penjumlahan dari 1 sampai %d adalah %d\n", n, total)
}
