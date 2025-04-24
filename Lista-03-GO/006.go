package main
import (
	"fmt"
)
func main() {
	var n int
	achou := false
	fmt.Scan(&n)

	for i := 0; i*(i+1)*(i+2) <= n; i++ {
		if i*(i+1)*(i+2) == n {
			achou = true
			break
		}
	} 
	if !achou {
		fmt.Printf("%d nao eh triangular\n", n)
	} else {
		fmt.Printf("%d eh triangular\n", n)
	}
}