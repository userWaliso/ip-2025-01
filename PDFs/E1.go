package main

import f "fmt"

func main() {
	var (
		x, cdd int
	)
	f.Scan(&x)
	var alt = make([]float64, x)

	for i := 0; i < x; i++ {
		f.Scan(&alt[i])
		cdd++
	}
	soma := 0.0
	for _, v := range alt {
		soma += v
	}
	media := soma / float64(cdd)

	for _, v := range alt {
		if v > media {
			f.Println(v)
		}
	}
}
