package main

import f "fmt"

func main() {
	var n, m int
	f.Scan(&n, &m)

	matriz := make([][]int, n)
	for i := range matriz {
		matriz[i] = make([]int, m)
		for j := range matriz[i] {
			f.Scan(&matriz[i][j])
		}
	}

	visitado := make([][]bool, n)
	for i := range visitado {
		visitado[i] = make([]bool, m)
	}

	
	dx := []int{0, 1, 0, -1}
	dy := []int{1, 0, -1, 0}
	dir := 0 

	x, y := 0, 0

	for i := 0; i < n*m; i++ {
		f.Printf("%d ", matriz[x][y])
		visitado[x][y] = true

		nx := x + dx[dir]
		ny := y + dy[dir]

		if nx < 0 || nx >= n || ny < 0 || ny >= m || visitado[nx][ny] {
			dir = (dir + 1) % 4 
			nx = x + dx[dir]
			ny = y + dy[dir]
		}
		x, y = nx, ny
	}
	f.Println()
}
