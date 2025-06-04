package main

import f "fmt"

func main() {
    var n int
    f.Scan(&n)

    matriz := make([][]int, n)
    for i := 0; i < n; i++ {
        matriz[i] = make([]int, n)
        for j := 0; j < n; j++ {
            f.Scan(&matriz[i][j])
        }
	}
	for i,linha := range matriz {
		for j := len(linha)-1; j >= 0; j-- {
			f.Printf("%d", matriz[j][i])
		}
		f.Println()
	}
}