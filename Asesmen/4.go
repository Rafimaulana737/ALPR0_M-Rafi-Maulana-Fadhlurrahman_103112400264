package main
import "fmt"

func main() {
	var a, b, c, d float64
	fmt.Println("Masukkan elemen-elemen matriks 2x2 (a b c d):")
	fmt.Scan(&a, &b, &c, &d)

	determinan := (a * d) - (b * c)
	fmt.Printf("Determinan: %.1f\n", determinan)
}