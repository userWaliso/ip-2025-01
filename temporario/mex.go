package main
import (
	f"fmt"
	"sort"
)

func main() {
	var n, k int

	f.Scan(&n, &k)
	var (
		vetor = make([]int, n)
		// submex[] int
)

	for i:=0;i<n;i++ {
		f.Scan(&vetor[i])
	}

	sort.Ints(vetor)
	// subvetor := vetor[0:k]
	copia:= vetor[k:len(vetor)]
	sort.Ints(copia)
	f.Println(copia[0])



	// vetor = vetor[k:len(vetor)]

	// for _,v := range vetor {	
	// 	if subvetor[2] != v {
	// 		submex = append(submex, v)
	// 	}
	// }
	// min:=submex[0]
	// for _,v := range submex {
	// 	if min > v {
	// 		min = v
	// 	}
	// }
	// f.Println(min)

}