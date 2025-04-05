package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Scan(&n)

	if 5 < n && n < 2000 {
		for i := 2; i <= n; i += 2 {
			fmt.Printf("%d ^2 = %d\n", i, i*i)
		}
	}
}
