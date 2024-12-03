package main

import "fmt"

func main() {
	var tebak int 
	
	for {
		var tebak int
		fmt.Print("Tebak angka (1-10): ")
		fmt.Scan(&tebak)
		if tebak == angkaRahasia {
			fmt.Println("Selamat, tebakan Anda benar!")
			break
		}
		fmt.Println("Tebakan Anda salah, coba lagi.")
	}
}
