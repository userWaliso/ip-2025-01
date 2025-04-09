package main

import (
	"fmt"
)

func main() {
	var ano int

	fmt.Scan(&ano)

	if ano%100 == 0 {
		if ano%4 == 0 && ano%400 == 0 {
			fmt.Print("Bissexto\n")
		} else {
			fmt.Print("Nao bissexto\n")
		}
	} else {
		if ano%4 == 0 {
			fmt.Print("Bissexto\n")
		} else {
			fmt.Print("Nao bissexto\n")
		}
	}
}
