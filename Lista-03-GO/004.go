package main
import (
	"fmt"
)
func main() {
	var n, i int
	

	for {
			fmt.Scan(&n)
			achou := false
		if n <=0 {
			return
		}	

		for i = 1; i*i <= n; i++ {
			if i*i == n {
				achou = true
				break
			 }	
		}
		if !achou {
				fmt.Printf("%d nao é quadrado perfeito\n", n)
			} else {
				fmt.Printf("%d é quadrado perfeito\n", n)
		} 
	}
}