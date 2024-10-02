package main

import (
    "fmt"
    "math"
)

func main() {
    var r float64
    const pi = 3.1415926535

    fmt.Print("Masukkan jari-jari bola: ")
    fmt.Scan(&r)

    volume := (4.0 / 3.0) * pi * math.Pow(r, 3)
    luas := 4 * pi * math.Pow(r, 2)

    fmt.Printf("Volume bola = %.4f\n", volume)
    fmt.Printf("Luas kulit bola = %.4f\n", luas)
}
