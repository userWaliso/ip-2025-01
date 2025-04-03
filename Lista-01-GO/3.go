package main
import "fmt"

func main(){
	var n1, n2, n3, valor float64

	fmt.Println("Escreva o número")
	fmt.Scan(&n1, &n2, &n3)
		if (n1 >= 10 && n1 == 0 && n2 >= 10 && n2 == 0 && n3 >= 10 && n3 == 0) {
		fmt.Println("DÍGITO INVÁLIDO")
			}else{
				valor = 100 * n1 + 10 * n2 + n3
					fmt.Println(valor)
						fmt.Println(valor * valor)
							fmt.Println("DÍGITO VÁLIDO")
	}
}