package main

import (
	"fmt"
)

func main() {
	var a, r, n, sn float64

	fmt.Scan(&a, &r, &n)

	sn = n * (2*a + (n-1)*r) / 2
	fmt.Println(sn)
}
