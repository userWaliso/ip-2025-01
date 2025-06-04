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

	soma := make([][]int, n)
	for i := 0; i < n; i++ {
		soma[i] = make([]int, m)
	}

	direcoes := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1},           {0, 1},
		{1, -1},  {1, 0},  {1, 1},
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			termo := 0
			for _,v := range direcoes {
				ni, nj := i+v[0], j+v[1]
				if ni >= 0 && ni < n && nj >= 0 && nj < m {
					termo += matriz[ni][nj]
				}
			}
			soma[i][j] = termo
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			f.Printf("%d ", soma[i][j])
		}
		f.Println()
	}
}
