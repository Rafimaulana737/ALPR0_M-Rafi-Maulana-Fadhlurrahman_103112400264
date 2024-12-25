package main
import "fmt"

func main() {
    var nama, nim, kelas string
    fmt.Scanln(&nama, &nim, &kelas)
    fmt.Printf("Perkenalkan, saya %s dari kelas %s dengan NIM %s.\n", nama, kelas, nim)
}
