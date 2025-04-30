package main
import (
	f"fmt"
	m"math"
)
func main() {
	var vetor []float64
	for i:=0;i<15;i++ {
		var x, raiz float64
		f.Scan(&x)
		if x < 0 {
			raiz = -1
		} else {
			raiz = m.Sqrt(x)
		}
		vetor = append(vetor, raiz)
	}
	f.Println(vetor)
}