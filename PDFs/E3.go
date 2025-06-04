package main

import f "fmt"

func main() {
	var x, v int
	f.Scan(&x)
	vetor := make([]int, x)

	for i := 0; i < x; i++ {
		f.Scan(&vetor[i])
	}
	if x == 1 {
		f.Println("0")
		return
	}

	var ivete []int

	min := vetor[1]
	max := vetor[len(vetor)-2]

	ivete = append(ivete, min)
	for i := 0; i < len(vetor); i++ {

		if i > 0 && i < len(vetor)-1 {

			v = vetor[i-1] + vetor[i+1]
			ivete = append(ivete, v)
		}
	}
	ivete = append(ivete, max)
	for _, n := range ivete {
		f.Printf("%d ", n)
	}
	f.Println()
}
