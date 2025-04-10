package main
import (
	"fmt"
)

func busca(lista []int, n int) int {
	for i, v := range lista {
		if v == n {
			return i
		}
	}
	return -1
}
func main() {
	var n int
	fmt.Scan(&n)

	lista := []int{10, 20, 30, 40, 50}

	vet := busca(lista, n)

		fmt.Println(vet)
	
}