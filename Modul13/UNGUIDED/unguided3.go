package main

import "fmt"

func main() {
	var g1, g2, g3, g4 string
	var done bool

	done = true
	
	for i := 1; i <= 5; i++ {
		fmt.Print("Percobaan", i, ": ")
		fmt.Scan(&g1, &g2, &g3, &g4)
		if g1 != "merah" && g2 != "kuning" && g3 != "hijau" && g4 != "ungu" {
			done = false
		}
	}
	fmt.Print("Berhasil: ", done)
}
