package main

import (
	"fmt"
)

func main() {
	var x, y, z int

	fmt.Scan(&x, &y, &z)

	if x+y <= z || x+z <= y || z+y <= x {
		fmt.Printf("Nao forma triangulo\n")
		return
	}

	if x == y && y == z {
		fmt.Println("Equilatero")
	} else if x == y || x == z || y == z {
		fmt.Println("Isosceles")
	} else {
		fmt.Println("Escaleno")
	}
}
