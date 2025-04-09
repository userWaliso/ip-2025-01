package main

import (
	"fmt"
)

func main() {
	var n1, n2, n3 float64
	var faltas, lim int
	fmt.Scan(&n1, &n2, &n3, &faltas)

	lim = 64 * 0.25
	media := (n1 + n2 + n3) / 3

	if faltas > lim {
		fmt.Print("Reprovado por falta\n")
	} else if media >= 7 {
		fmt.Print("Aprovado\n")
	} else if 5 <= media && media < 7 {
		fmt.Print("Prova final\n")
	} else if media < 5 {
		fmt.Print("Reprovado\n")
	}
}
