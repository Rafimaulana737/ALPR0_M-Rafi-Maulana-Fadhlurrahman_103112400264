package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan jumlah bilangan: ")
	fmt.Scan(&n)

	bilangan := make([]int, n)
	fmt.Println("Masukkan bilangan-bilangan:")
	for i := 0; i < n; i++ {
		fmt.Scan(&bilangan[i])
	}

	fmt.Println("Hasil perkalian:")
	for i := 0; i < n; i++ {
		fmt.Print(bilangan[i]*(i+1), " ")
	}
	fmt.Println()
}
