package main
import f "fmt"

func main() {
	var x int
	f.Scan(&x)
	vetor := make([]int, x)

	for i:=0;i<x;i++ {
		f.Scan(&vetor[i])
	}
	soma := 0
	for _,v := range vetor {
		soma += v
		f.Printf("%d ",soma)
	}
	f.Println()
}