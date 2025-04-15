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
	var x, rs float64
	fmt.Scan(&x)
	rs=x
	for i:=1;i<=20;i++ {
		pot := math.Pow(x, float64(i))
		fat := float64(fatorial(i))
		termo := pot / fat
		if i%2 == 0 {
			rs += termo
		} else {
			rs -= termo
		}

	}
	fmt.Print(rs)
}