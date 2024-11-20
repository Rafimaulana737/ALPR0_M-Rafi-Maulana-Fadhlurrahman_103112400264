package main

import "fmt"

func main() {

	var kode int
    fmt.Print("Masukkan kode menu (1-5): ")
    fmt.Scan(&kode)


    fmt.Println("Menu Restoran:")
    fmt.Println("1. Burger - Rp25,000")
    fmt.Println("2. Fried Chicken - Rp20,000")
    fmt.Println("3. French Fries - Rp15,000")
    fmt.Println("4. Soft Drink - Rp10,000")
    fmt.Println("5. Coffee - Rp15,000")


    switch kode {
    case 1:
        fmt.Println("Burger - Rp25,000")
    case 2:
        fmt.Println("Fried Chicken - Rp20,000")
    case 3:
        fmt.Println("French Fries - Rp15,000")
    case 4:
        fmt.Println("Soft Drink - Rp10,000")
    case 5:
        fmt.Println("Coffee - Rp15,000")
    default:
        fmt.Println("Kode menu tidak valid")
    }
}
