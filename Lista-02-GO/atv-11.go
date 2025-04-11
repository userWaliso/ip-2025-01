package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Scan(&x)

	f := 8 / (2-x)

	fmt.Println(f)
}