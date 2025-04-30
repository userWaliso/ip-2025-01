package main 
import f"fmt"
func main() {
	var x, v int
	f.Scan(&x)
	vetor := make([]int, x)

	for i:=0;i<x;i++ {
		f.Scan(&vetor[i])
	}
	var ivete[] int
	min:= vetor[1]
	max := vetor[len(vetor) - 2]
	for i:=0;i<len(vetor);i++ {
		if i != 0 {
			v = vetor[i-1] + vetor[i+1]
		} else if i == len(vetor) - 1 {
			
		}
		ivete = append(ivete, v)
	}
	println(min, max, ivete)
	
}