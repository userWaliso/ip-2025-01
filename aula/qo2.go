package main

import "fmt"

func main() {

	var n int
	fmt.Println("Escreva o número: ")
	fmt.Scan(&n)

	if n >= 20 && n <= 90{
		fmt.Println("Sim")
	}else {
		fmt.Println("Não")
	}


}