package main

import (
	"fmt"
)

func main() {
	var idd int

	fmt.Scan(&idd)

	if idd > 0 && idd < 16 {
		fmt.Println("Não eleitor\n")
	} else if idd >= 18 && idd <= 65 {
		fmt.Println("Eleitor Obrigatório\n")
	} else if idd >= 16 && idd < 18 || idd > 65 {
		fmt.Println("Eleitor Facultativo\n")
	}
}
