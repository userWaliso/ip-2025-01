package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, x1, x2, delta float64

	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	delta = math.Pow(b, 2) - 4*a*c
	x1 = (-b + math.Sqrt(delta)) / (2 * a)
	x2 = (-b - math.Sqrt(delta)) / (2 * a)

	if a == 0 {
		fmt.Printf("Nao e equacao do segundo grau\n")
		} else { if delta < 0 {
			fmt.Println("RAIZES IMAGINARIAS")
			} else if delta == 0 {
				fmt.Println("RAIZES IGUAIS")
				x1 = x2 
				fmt.Printf("%.5f\n", x1)
		
			} else {
				fmt.Println("RAIZES DISTINTAS")
				fmt.Printf("%.5f %.5f\n", x1, x2)
			}
	}
}
