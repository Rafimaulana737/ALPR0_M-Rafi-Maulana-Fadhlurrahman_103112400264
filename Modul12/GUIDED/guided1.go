package main

import "fmt"

func main(){
	var word string
	var repitation int
	
	fmt.Scan(&word, repitation)
	counter := 0

	for done := false; !done; {
		fmt.Println(word)
		counter++
		done = (counter >= repitation)
	}

}