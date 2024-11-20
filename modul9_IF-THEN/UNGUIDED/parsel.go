package main

import "fmt"

func main() {
    var beratGram, kg, sisaGram, biaya int
    fmt.Print("Berat parsel (gram): ")
    fmt.Scan(&beratGram)

    kg = beratGram / 1000
    sisaGram = beratGram % 1000

    biaya = kg * 10000
    if kg > 10 {
        fmt.Printf("Detail berat: %d kg + %d gr\nDetail biaya: Rp. %d\nTotal biaya: Rp. %d\n", kg, sisaGram, biaya, biaya)
    } else if sisaGram >= 500 {
        biaya += sisaGram * 5
    } else {
        biaya += sisaGram * 15
    }

    fmt.Printf("Detail berat: %d kg + %d gr\nDetail biaya: Rp. %d\nTotal biaya: Rp. %d\n", kg, sisaGram, biaya-kg*10000, biaya)
}
