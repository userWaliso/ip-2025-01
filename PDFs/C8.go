package main 
import (
	"fmt"
	"math"
)
func fatorial(a int) int {
	var fat int = 1 
	for i:=1; i<=a; i++{
		fat *= i
	}
	return fat
}

func main() {
	var (
	 x  float64
	)
	
	fmt.Scan(&x)
	
	for a := 0.0; a <= x+1e-9; a += 0.1 {
		sin := a
		sin -= math.Pow(a, 3)/float64(fatorial(3))
		sin += math.Pow(a, 5)/float64(fatorial(5))
		sin -= math.Pow(a, 7)/float64(fatorial(7))
		fmt.Printf("%.1f %.4f\n", a, sin)
	}
}

