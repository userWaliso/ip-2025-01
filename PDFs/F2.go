package main

import f "fmt"

func main() {
    var n, m int
    f.Scan(&n, &m)

    matriz := make([][]int, n)
    for i := 0; i < n; i++ {
        matriz[i] = make([]int, m)
        for j := 0; j < m; j++ {
            f.Scan(&matriz[i][j])
        }
	}

	for i := range matriz {
		somaL := 0
		for j := range matriz[i] {
			somaL += matriz[i][j]
		}
		f.Printf("%d ", somaL)
	}

	f.Println()
		
	for i := 0; i < m; i++ {
		somaC := 0
		for j := 0; j < n; j++ {
			somaC += matriz[j][i]
		}
		f.Printf("%d ", somaC)
	}
	f.Println()

}