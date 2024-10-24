package main

import "fmt"

func main() {
    var base, exponent, result int
    fmt.Print("Masukkan bilangan dasar dan pangkat: ")
    fmt.Scan(&base, &exponent)

    result = 1
    for i := 0; i < exponent; i++ {
        result *= base
    }

    fmt.Printf("Hasil %d dipangkatkan %d adalah %d\n", base, exponent, result)
}
