package main

import "fmt"

func main() {
	password := "1234"
	attempts := 3

	for attempts > 0 {
		var input string
		fmt.Print("Masukkan password: ")
		fmt.Scan(&input)

		if input == password {
			fmt.Println("Login berhasil!")
			return
		}
		attempts--
		if attempts > 0 {
			fmt.Println("Password salah, coba lagi.")
		}
	}
	fmt.Println("Login ditolak.")
}
