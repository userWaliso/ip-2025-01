package main
import (
	"fmt"
)
func main() {
	var n, a, i, soma int

	fmt.Scan(&n,&a)
	soma = 1
	for i = 1; i <= a; i++ {
		soma *= n 
	}
	fmt.Printf("%d\n", soma)
}