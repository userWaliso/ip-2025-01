package main

import f "fmt"

func main() {
    var n, k, m int
    f.Scan(&n, &k, &m)

    A := make([][]int, n)
	B := make([][]int, k)
    for i := 0; i < n; i++ {
        A[i] = make([]int, k)
        for j := 0; j < k; j++ {
            f.Scan(&A[i][j])
        }
	}
	for i := 0; i < k; i++ {
		B[i] = make([]int, m)
		for j := 0; j < m; j++ {
			f.Scan(&B[i][j])
		}
	}

	C := make([][]int, n)
    for i := 0; i < n; i++ {
        C[i] = make([]int, m)
        for j := 0; j < m; j++ {
            soma := 0
            for l := 0; l < k; l++ {
                soma += A[i][l] * B[l][j]
            }
            C[i][j] = soma
        }
    }
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			f.Printf("%d ", C[i][j])
		}
		f.Println()
	}
}