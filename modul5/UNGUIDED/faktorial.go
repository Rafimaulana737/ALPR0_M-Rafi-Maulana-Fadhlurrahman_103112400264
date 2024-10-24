package main

import "fmt"

func main() {
    var n, result int
    fmt.Print("Masukkan bilangan non-negatif: ")
    fmt.Scan(&n)

    result = 1
    for i := 2; i <= n; i++ {
        result *= i
    }

    fmt.Printf("Hasil faktorial dari %d adalah %d\n", n, result)
}
