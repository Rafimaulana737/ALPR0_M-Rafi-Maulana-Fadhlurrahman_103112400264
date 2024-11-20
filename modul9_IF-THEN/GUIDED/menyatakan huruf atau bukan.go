package main

import "fmt"

func main() {
	var huruf string

	fmt.Print("Masukkan huruf: ")
	fmt.Scan(&huruf)

	// Menggunakan kondisi untuk mengecek apakah input adalah huruf vokal
	if huruf == "A" || huruf == "I" || huruf == "U" || huruf == "E" || huruf == "O" || 
	   huruf == "a" || huruf == "i" || huruf == "u" || huruf == "e" || huruf == "o" {
		fmt.Println("Huruf vokal")
	} else {
		fmt.Println("Huruf konsonan")
	}
}
