package main
import f "fmt"

func main() {
	var x int
	f.Scan(&x)
	vetor := make([]int, x)

	for i:=0;i<x;i++ {
		f.Scan(&vetor[i])
	}

	comp := make(map[int]struct{})
    for _, v := range vetor {
       comp[v] = struct{}{} 
	}
	f.Println(len(comp))
}