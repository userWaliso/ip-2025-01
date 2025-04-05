package main

import (
	"fmt"
	"math"
)

func main() {
	var h, a, ab, v float64

	fmt.Scan(&h)
	fmt.Scan(&a)

	ab = 3 * math.Pow(a, 2) * math.Sqrt(3) / 2
	v = ab * h / 3

	fmt.Printf("O VOLUME DA PIRAMIDE É = %.2f METROS CÚBICOS\n", v)

}
