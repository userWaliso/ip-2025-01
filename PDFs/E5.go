package main 
import f "fmt"

func main() {
	var x, count int
	f.Scan(&x)
	var vetor = make([]int, x)
	var vishe[] int

	for i:=0;i<x;i++ {
		f.Scan(&vetor[i])
	}
	for i:=0;i<x-1;i++ {
		if vetor[i+1] == vetor[i] + 1 {
			count++
			vishe = append(vishe, count)
		} else if vetor[i+1] != vetor[i] + 1 {
			count = 1
		}
	}
	f.Println(vishe)
}