package main

import "fmt"

func main() {
	var input string

	for {
		fmt.Print("Masukkan kata: ")
		fmt.Scan(&input) // Membaca input satu kata
		if input == "telkom" {
			fmt.Println("Program selesai.")
			break
		}
		fmt.Println("Anda mengetik:", input)
	}
}
