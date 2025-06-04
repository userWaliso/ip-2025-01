package main

import f "fmt"

func main() {
    var n, m int
    f.Scan(&n, &m)

    matriz := make([][]float64, n)
    for i := 0; i < n; i++ {
        matriz[i] = make([]float64, m)
        for j := 0; j < m; j++ {
            f.Scan(&matriz[i][j])
        }
	}
	rapaz := make([]float64, 0, n)
	for i := 0; i < n; i++ {
		max := matriz[i][0]
		for j := 0; j < m; j++ {
			if max < matriz[i][j] {
				max = matriz[i][j]
				} 
			}
			rapaz = append(rapaz, max)
		}

		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				f.Printf("%.6f ", matriz[i][j]/rapaz[i])
			}
			f.Println()
		// f.Println(rapaz)
	}
}