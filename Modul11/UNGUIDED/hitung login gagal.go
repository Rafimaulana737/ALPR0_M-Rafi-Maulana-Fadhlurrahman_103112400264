package main

import "fmt"

func main() {
	var username, password string
	var attempts int

	for {
		fmt.Print("Masukkan username: ")
		fmt.Scan(&username)
		fmt.Print("Masukkan password: ")
		fmt.Scan(&password)

		if username == "Admin" && password == "Admin" {
			break
		}
		attempts++
		fmt.Println("Login gagal, coba lagi.")
	}

	fmt.Printf("%d percobaan gagal login\n", attempts)
}
