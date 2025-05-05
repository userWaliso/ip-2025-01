package main
import f"fmt"
func main() {
	var (
		x int
		norm float64
	)
	f.Scan(&x)
	var vetor = make([]float64, x)

	for i:=0;i<x;i++ {
		f.Scan(&vetor[i])
	}
	min := vetor[0]
	max := vetor[0]
	for _,v := range vetor {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	
	for _,v := range vetor {
		if max - min == 0 {
			norm = 0
		} else {
			norm = (v - min) / (max - min)
		}
		f.Printf("%.2f ", norm)
	}
}