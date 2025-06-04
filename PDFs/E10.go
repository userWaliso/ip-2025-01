package main

import f "fmt"

func main() {
	var x int
	f.Scan(&x)
	vetor := make([]int, x)
	var mentira bool = true

	for i := 0; i < x; i++ {
		f.Scan(&vetor[i])
	}
	for i := 0; i < x/2; i++ {
		if vetor[i] != vetor[x-1-i] {
			mentira = false
			break
		}
	}

	if mentira {
		f.Print("SIM\n")
	} else {
		f.Print("NAO\n")
	}
}
