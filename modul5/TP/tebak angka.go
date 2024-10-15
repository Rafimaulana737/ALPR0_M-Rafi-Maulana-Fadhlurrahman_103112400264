package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    target := rand.Intn(100) + 1
    var guess int
    attempts := 5

    fmt.Println("Permainan Tebak Angka (1-100)")
    for i := 0; i < attempts; i++ {
        fmt.Printf("Percobaan %d: Masukkan tebakan kamu: ", i+1)
        fmt.Scan(&guess)

        if guess == target {
            fmt.Println("Selamat! Tebakan kamu benar.")
            return
        } else if guess < target {
            fmt.Println("Tebakan terlalu kecil.")
        } else {
            fmt.Println("Tebakan terlalu besar.")
        }
    }

    fmt.Println("Maaf, kamu kehabisan percobaan. Angka yang benar adalah:", target)
}
