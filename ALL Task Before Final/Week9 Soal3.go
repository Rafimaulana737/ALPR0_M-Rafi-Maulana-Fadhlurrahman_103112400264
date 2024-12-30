package main

import "fmt"

func main() {
	var char rune
	fmt.Print("Masukkan satu karakter: ")
	fmt.Scanf("%c", &char)

	if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' {
		if char != 'a' && char != 'e' && char != 'i' && char != 'o' && char != 'u' && char != 'A' && char != 'E' && char != 'I' && char != 'O' && char != 'U' {
			fmt.Println("konsonan")
			return
		}
	}
	fmt.Println("bukan konsonan")
}
