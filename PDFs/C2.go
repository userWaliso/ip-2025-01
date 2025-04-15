package main
import (
	"fmt"
)
func main() {
	var n, a, i, soma, cdd int

	fmt.Scan(&n,&a)
	for i = n; i <= a; i++ {
		if i%2 == 0{
			soma += i
			cdd ++
		}
	}
	if cdd > 0 {
		media := soma / cdd
	fmt.Printf("%d\n%d\n",soma, media)
	} else {
		media := 0
		fmt.Printf("%d\n%d\n",soma, media)
	}
}