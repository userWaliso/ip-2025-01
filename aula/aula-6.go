package main

import (
	f "fmt"
)

func buscabinaria(x int) int {
	array := []int{2, 3, 7, 12, 15, 17, 23, 30, 40, 50}

	e := 0
	d := len(array) - 1

	for e <= d {
		m := (e + d) / 2
		if array[m] == x {
			return m
		}
		if array[m] < x {
			e = m + 1
		}
		if array[m] > x {
			d = m - 1
		}
	}
	return -1
}

func main() {

	f.Println(buscabinaria(50))
}
