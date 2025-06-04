package main
import (
	f"fmt"
)
func main() {
	var (
		num = []int{}
		indices = []int{}
	)

	for i:=0;i<10;i++ {
		var x int
		f.Scan(&x)
		num = append(num, x)
	}
	for i, v := range num {
		if v > 50 {
			indices = append(indices, i)
		}
	}
	f.Printf("Números maiores que 50: ")
	for _, v := range num {
		if v > 50 {
			f.Printf("%d ", v)
		}
	}
	f.Printf("\nRespectivos índices: %d", indices)
}