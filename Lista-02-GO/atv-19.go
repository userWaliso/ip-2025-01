package main
import (
	"fmt"
	"math"
)
func main() {
	var esc int
	const pi = 3.14

	fmt.Scan(&esc)
	if esc < 1 || esc > 3 {
		fmt.Println("NUMERO INVALIDO")
		return
	}
	if esc == 1 {
		var r, h float64

		fmt.Println("Dẽ o valor do raio e da altura")
		fmt.Scan(&r,&h)

		volum := (pi * math.Pow(r, 2)* h)/3
		area := pi * r * math.Sqrt(r * r + h * h)
		fmt.Printf("Volume = %.2f Area = %.2f", volum, area)
	} else if esc == 2 {
		var r, h float64

		fmt.Println("Dẽ o valor do raio e da altura")
		fmt.Scan(&r,&h)

		volum := (pi * math.Pow(r, 2)* h)
		area := 2 * pi * r * h
		fmt.Printf("Volume = %.2f Area = %.2f", volum, area)
	} else if esc == 3 {
		var r float64

		fmt.Println("Dẽ o valor do raio e da altura")
		fmt.Scan(&r)

		volum := (pi* math.Pow(r,3))/3
		area := 4*pi*math.Pow(r,2)
		fmt.Printf("Volume = %.2f Area = %.2f", volum, area)
	}
}