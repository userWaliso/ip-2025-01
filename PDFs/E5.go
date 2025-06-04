package main

import f "fmt"

func main() {
	var x int
	f.Scan(&x)
	var vetor = make([]int, x)
	var vishe []int
	count := 1

	for i := 0; i < x; i++ {
		f.Scan(&vetor[i])
	}
	for i := 0; i < x-1; i++ {
		if vetor[i+1] == vetor[i]+1 {
			count++
			vishe = append(vishe, count)
		} else {
			count = 1
		}
	}
	if count >= 1 {
		vishe = append(vishe, count)
	}

	max := vishe[0]
	for _, v := range vishe {
		if max < v {
			max = v
		}
	}
	f.Println(max)
}
