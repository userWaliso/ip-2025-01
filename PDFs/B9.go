package main

import (
	"fmt"
)

func main() {
	var alt, peso, imc float64
	fmt.Scan(&peso, &alt)

	imc = peso / (alt * alt)

	if imc < 18.5 {
		fmt.Print("Abaixo do peso\n")
	} else if 18.5 <= imc && imc < 25 {
		fmt.Print("Peso normal\n")
	} else if 25 <= imc && imc < 30 {
		fmt.Print("Sobrepeso\n")
	} else if imc >= 30 {
		fmt.Print("Obesidade\n")
	}
}
