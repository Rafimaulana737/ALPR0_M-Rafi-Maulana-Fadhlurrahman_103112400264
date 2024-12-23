package main

import "fmt"

func main() {
	var bil int

	fmt.Print("Masukan bilangan bulat: ")
	fmt.Scan(&bil)

	if bil <= 1 {
		fmt.Println("Bukan Prima")
		return
	}

	for i := 2; i < bil; i++ {
		if bil%i == 0 {
			fmt.Println("Bukan Prima")
			return
		}
	}
	fmt.Println("Prima")
}
