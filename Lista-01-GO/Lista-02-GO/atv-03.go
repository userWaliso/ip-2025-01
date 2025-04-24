
package main

import "fmt"

func main() {
	var a, b int
	
	fmt.Scan(&a, &b)
	c := a + b 
	if c > 20 {
		fmt.Printf("%d\n",c + 8)
	} else {
		fmt.Printf("%d\n",c - 5)
	}
}
