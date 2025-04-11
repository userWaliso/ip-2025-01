package main

import (
	"fmt"
)
func main() {
var dia,mes, ano int

	fmt.Scan(&dia, &mes, &ano)

	switch mes {
		case 1: fmt.Printf("%d de Janeiro de %d\n", dia, ano)
		case 2: fmt.Printf("%d de Fevereiro de %d\n", dia, ano)
		case 3: fmt.Printf("%d de Março de %d\n", dia, ano)
		case 4: fmt.Printf("%d de Abril de %d\n", dia, ano)
		case 5: fmt.Printf("%d de Maio de %d\n", dia, ano)
		case 6: fmt.Printf("%d de Junho de %d\n", dia, ano)
		case 7: fmt.Printf("%d de Julho de %d\n", dia, ano)
		case 8: fmt.Printf("%d de Agosto de %d\n", dia, ano)
		case 9: fmt.Printf("%d de Setembro de %d\n", dia, ano)
		case 10: fmt.Printf("%d de Outubro de %d\n", dia, ano)
		case 11: fmt.Printf("%d de Novembro de %d\n", dia, ano)
		case 12: fmt.Printf("%d de Dezembro de %d\n", dia, ano)
	}
}