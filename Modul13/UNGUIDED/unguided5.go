package main

import "fmt"

func main() {
	var kantong1, kantong2 float64

	for {
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&kantong1, &kantong2)

		if kantong1 < 0 || kantong2 < 0 {
			fmt.Println("Proses Selesai")
			break
		}
		if kantong1+kantong2 > 150 {
			fmt.Println("Program Selesai")
			break
		}
		if kantong1-kantong2 > 9 || kantong2-kantong1 > 9 {
			fmt.Println("Sepeda motor pak Andi akan oleng: true")
		} else {
			fmt.Println("Sepeda motor pak Andi akan oleng: false")
		}
	}

}
