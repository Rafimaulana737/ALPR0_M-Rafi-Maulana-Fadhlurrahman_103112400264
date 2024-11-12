package main

import "fmt"

func main(){
	var i int
	fmt.Print("masukan bilangan bulat")
	fmt.Scan(&i)

	if i < 0 {
		fmt.Println(true)
	}else{
		fmt.Println(false)
	}
}