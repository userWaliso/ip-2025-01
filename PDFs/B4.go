package main

import (
	"fmt"
)

func main() {
	var idd int

	fmt.Scan(&idd)

	if idd >= 5 && idd <= 7 {
		fmt.Printf("Infantil A\n")
	} else if idd >= 8 && idd <= 10 {
		fmt.Printf("Infantil B\n")
	} else if idd >= 12 && idd <= 13 {
		fmt.Printf("Juvenil A\n")
	} else if idd >= 14 && idd <= 17 {
		fmt.Printf("Juvenil B")
	} else if idd >= 18 {
		fmt.Printf("Adulto\n")
	}
}
