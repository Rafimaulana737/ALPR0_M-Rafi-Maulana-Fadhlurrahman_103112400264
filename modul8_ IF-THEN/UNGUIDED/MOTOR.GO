package main
import "fmt"

func main() {
	
    var jumlahOrang, jumlahMotor int
	
    fmt.Scan(&jumlahOrang)
    
	jumlahMotor = (jumlahOrang + 1) / 2
    
	fmt.Println(jumlahMotor)
}
