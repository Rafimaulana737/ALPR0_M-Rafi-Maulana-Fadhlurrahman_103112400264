package main

import "fmt"

func main() {
	var n, count int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&n)
	for n > 0 {
		count++
		n /= 10
	}
	fmt.Println("Jumlah digit:", count)
}
