package main
import (
	"fmt"
)
func main(){
	var dia, preco, preco1, preco2, valor float64
	var lan int

	fmt.Println("Preço do produto e dia da semana")
	fmt.Scan(&preco, &dia)
	fmt.Println("O produto é de lançamento? 1 para sim e 2 para não")
	fmt.Scan(&lan)

	if dia == 2 || dia == 3 || dia == 5 {
		preco1 = preco * 0.60
	}
	if lan == 1 {
		preco2 = (preco * 1.15) - preco
	}
	valor = preco1 + preco2
	fmt.Println(valor)
}