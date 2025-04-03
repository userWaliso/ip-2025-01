package main
import "fmt"

func main () {
	var n1, n2, n3, n4 float64

	fmt.Println("Escreva o número: ")
	fmt.Scan(&n1, &n2, &n3)

	n4 = (n1 + n2 + n3)/3

	fmt.Println("A média é: ", n4,"\n")

	if (n4 > 6) {
		fmt.Println("APROVADO")
	}else {
		fmt.Println("REPROVADO")
	}
}