package main

import (
	"fmt"
)

func main() {
	var a, b int

	fmt.Scan(&a, &b)

	if a % b == 0 {
		fmt.Printf("%d E DIVISIVEL POR %d\n", a, b)
	} else {
		fmt.Printf("A NAO É DIVISIVEL POR B\n")
	} 
}
