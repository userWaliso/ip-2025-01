package main

import (
	"fmt"
	"math"
)

func main() {
	var n float64

	fmt.Scan(&n)

	if n >= 0 {
		n = (math.Sqrt(float64(n)))
	} else if n <= 0 { 
		n = (math.Pow(float64(n), 2))
	}
	fmt.Printf("%.2f\n", n)
}