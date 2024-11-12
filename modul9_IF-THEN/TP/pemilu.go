package main
import "fmt"

func main() {
    var umur int
    var kewarganegaraan string

    fmt.Print("Masukkan umur: ")
    fmt.Scan(&umur)
    fmt.Print("Masukkan kewarganegaraan (WNI/WNA): ")
    fmt.Scan(&kewarganegaraan)

    if umur >= 17 && kewarganegaraan == "WNI" {
        fmt.Println("Anda bisa mengikuti pemilu")
    } else {
        if umur < 17 {
            fmt.Println("Tidak memenuhi syarat: usia kurang dari 17 tahun")
        }
        if kewarganegaraan != "WNI" {
            fmt.Println("Tidak memenuhi syarat: bukan WNI")
        }
    }
}
