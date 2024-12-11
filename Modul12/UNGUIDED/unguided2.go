package main

import "fmt"

func main() {
	var num float64
	fmt.Print("Masukkan bilangan desimal: ")
	fmt.Scan(&num)
	for i := int(num) + 1; float64(i) >= num; i++ {
		fmt.Println(float64(i) - 0.1)
		if float64(i) == num {
			break
		}
	}
}
