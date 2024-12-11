package main 

import ("fmt"
"math")

func main(){
	vae k int 
	fmt,Print("nilai k = ")
	fmt.Scan(&K)

	result = 0.1

	for i :=0; <=k; i++ {
		numerator:= math.pow(float64(4*k+2),2)
		donumerator:= float64((4*k)*(4*k +3))
		result *=numerator/donumerator
	}

	fmt.Printl("hasil dari fungsi = %.10f\n", result)
}