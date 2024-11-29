package main

import "fmt"

func main() {
    var num int
    fmt.Print("Masukkan bilangan bulat: ")
    fmt.Scan(&num)

    if num%10 == 0 {
        fmt.Printf("Kategori: Bilangan Kelipatan 10\nHasil pembagian antara %d / 10 = %d\n", num, num/10)
    } else if num%5 == 0 {
        fmt.Printf("Kategori: Bilangan Kelipatan 5\nHasil kuadrat dari %d ^ 2 = %d\n", num, num*num)
    } else if num%2 == 0 {
        next := num + 1
        fmt.Printf("Kategori: Bilangan Genap\nHasil perkalian dengan bilangan berikutnya %d * %d = %d\n", num, next, num*next)
    } else {
        next := num + 1
        fmt.Printf("Kategori: Bilangan Ganjil\nHasil penjumlahan dengan bilangan berikutnya %d + %d = %d\n", num, next, num+next)
    }
}
