package main

import "fmt"

func main(){
	var i int

	fmt.Print("masukan bilanagan bulat")
	fmt.Scan(&i)
	if i > 1 {
		fmt.Print("positif")

	} else{
		fmt.Print("negatif")
	}
}