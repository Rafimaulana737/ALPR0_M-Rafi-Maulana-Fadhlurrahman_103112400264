package main

import "fmt"

func main() {
	var n, sum int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&n)
	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	if sum == n {
		fmt.Println("Ya")
	} else {
		fmt.Println("Tidak")
	}
}
