package main

import "fmt"

func main(){
	var i int

	fmt.Println("masukan baris bilangan:")
	fmt.Scan(&n)
	fmt.Println("masukan bilangan akhir")
	fmt.Scanl(&n)

	for i := n;<= n; i ++{
		fmt.println(i)
	}
