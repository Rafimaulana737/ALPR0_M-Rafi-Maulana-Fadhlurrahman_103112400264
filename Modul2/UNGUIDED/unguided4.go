package main
import "fmt"

func main() {
    var f, c float64
    fmt.Scan(&f)
    c = (f - 32) * 5 / 9
    fmt.Printf("Suhu dalam Celsius: %.2f\n", c)
}
