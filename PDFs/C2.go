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

package main
import (
	f"fmt"
)
func buscabinaria() {
	var x
	array := make([10], int) {2, 3, 7, 12, 15, 17, 23, 30, 40, 50}

	e := 0
	d := len[array] - 1

	for e <= D {
		m:= (e + d) / 2
		if len[array] == x
			return m
		if len[array] < x
			e = m + 1
		if len[array] > x
			d = m - 1
	} return -1

}
