package main

import "fmt"

func main() {
	var bil, hasil int

	fmt.Print("Masukan sebuah bilangan: ")
	fmt.Scan(&bil)

	for i := 1; i <= bil; i++ {
		if i%2 != 0 {
			hasil++
		}
	}
	fmt.Println("Terdapat", hasil, "bilangan ganjil")
}
