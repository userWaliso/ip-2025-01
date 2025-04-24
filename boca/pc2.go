package main
import (
	"fmt"
)
func main() {
	var x int
	var conceito string
	fmt.Scan(&x)
	if x < 0 && x > 10 {
		return
	}
	if x < 3 {
		conceito = "E"
	} else if x >= 3 && x < 5 {
		conceito = "D"
	} else if x >= 5 && x < 7 {
		conceito = "C"
	} else if x >= 7 && x < 9 {
		conceito = "B"
	} else {
		conceito = "A"
	}
	fmt.Println(conceito)
}