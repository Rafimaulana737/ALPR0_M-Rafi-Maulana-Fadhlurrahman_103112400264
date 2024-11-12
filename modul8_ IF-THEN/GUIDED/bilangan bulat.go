package main 

import "fmt"

func main(){
	var i int
	fmt.Print("masukan bilangan")
	fmt.Scan (&i)
	if i <0 {
		i = -i

	}
	fmt.Print(i)
	
}