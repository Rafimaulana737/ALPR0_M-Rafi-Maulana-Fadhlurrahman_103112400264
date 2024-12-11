package main

import "fmt"

func main() {
	var target, total, donasi, count int
	fmt.Print("Masukkan target donasi: ")
	fmt.Scan(&target)
	for total < target {
		count++
		fmt.Printf("Masukkan donasi ke-%d: ", count)
		fmt.Scan(&donasi)
		total += donasi
		fmt.Printf("Total terkumpul: %d\n", total)
	}
	fmt.Printf("Target tercapai! Total donasi: %d dari %d donatur.\n", total, count)
}
