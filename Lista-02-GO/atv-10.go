package main 

import (
	"fmt"
)

func main() {
	var n, volta, norte, sul, nordeste, centro int
	
	fmt.Scan(&n)
	if n < 1 || n > 5 {
		fmt.Print("Número inválido")
		return 
	}
	fmt.Print("A viagem terá volta?\n")
	fmt.Scan(&volta)

	if n == 1 && volta == 1{ 
		norte = 900
		fmt.Printf("o valor da ida e volta é R$:%d.00\n", norte) 
		} else if n == 1 && volta == 2 {
			norte = 500
			fmt.Printf("o valor da ida é R$:%d.00\n", norte)
	}
	if n == 2 && volta == 1{ 
		nordeste = 650
		fmt.Printf("o valor da ida e volta é R$:%d.00\n", nordeste) 
		} else if n == 2 && volta == 2 {
			nordeste = 350
			fmt.Printf("o valor da ida é R$:%d.00\n", nordeste)
	}
	if n == 3 && volta == 1{ 
		centro = 600
		fmt.Printf("o valor da ida e volta é R$:%d.00\n", centro) 
		} else if n == 3 && volta == 2 {
			centro = 350
			fmt.Printf("o valor da ida é R$:%d.00\n", centro)
	}
	if n == 4 && volta == 1{ 
		sul = 550
		fmt.Printf("o valor da ida e volta é R$:%d.00\n", sul) 
		} else if n == 4 && volta == 2 {
			sul = 300
			fmt.Printf("o valor da ida é R$:%d.00\n", sul)
	}
}