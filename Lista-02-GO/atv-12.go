package main

import (
	"fmt"
)

func main() {
	var idd int

	fmt.Scan(&idd)

	if idd >= 0 && idd <= 2 {
		fmt.Printf("Recém nascido\n")
	} else if idd >= 3 && idd <= 11 {
		fmt.Printf("Criança\n")
	} else if idd >= 12 && idd <= 19 {
		fmt.Printf("Adolescente\n")
	} else if idd >= 20 && idd <= 55 {
		fmt.Printf("Adulto\n")
	} else if idd > 55 {
		fmt.Printf("Idoso\n")
	}
}
