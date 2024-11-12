package main
import "fmt"

func main() {
    var huruf string
    fmt.Print("Masukkan huruf: ")
    fmt.Scan(&huruf)

    switch huruf {
    case "A", "I", "U", "E", "O", "a", "i", "u", "e", "o":
        fmt.Println("Huruf Vokal")
    default:
        fmt.Println("Huruf Konsonan")
    }
}
