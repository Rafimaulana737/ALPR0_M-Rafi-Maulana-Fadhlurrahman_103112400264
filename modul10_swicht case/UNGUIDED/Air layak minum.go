package main

import "fmt"

func main() {
    var ph float64
    fmt.Print("Masukkan nilai pH: ")
    fmt.Scan(&ph)

    if ph < 0 || ph > 14 {
        fmt.Println("Input tidak valid, rentang pH 0 - 14")
    } else if ph >= 6.5 && ph <= 8.6 {
        fmt.Println("Air layak minum")
    } else {
        fmt.Println("Air tidak layak minum")
    }
}
