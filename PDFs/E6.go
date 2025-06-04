package main
import f "fmt"

func main() {
	var (
		x, sort int
	)
	f.Scan(&x)
	vetor := make([]int, x)

	for i:=0;i<x;i++ {
		f.Scan(&vetor[i])
	}
	for i:=len(vetor)-1;i >= 0;i-- {
		for j:=1; j<=i;j++ {
			if vetor[j-1] > vetor[j] {
				sort = vetor[j-1]
				vetor[j-1] = vetor[j]
				vetor[j] = sort
			}
		}
	}
	for _,v := range vetor {
		f.Printf("%d ", v)
	}
	
	f.Println()
}