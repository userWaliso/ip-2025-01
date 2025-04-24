package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Scan(&n)
	if n%2 == 0 {
		fmt.Printf("%d Par\n", n)
	} else {
		fmt.Printf("%d Impar\n", n)
	}
}
