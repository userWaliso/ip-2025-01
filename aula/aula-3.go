package main
import (
	"fmt"
)
func soma(a, b int) int {
	return a + b
}
func main() {
	var a, b int

	fmt.Scan(&a, &b)

	x := soma(a, b)
	fmt.Println(x)
}