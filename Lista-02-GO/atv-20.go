package main
import (
	"fmt"
)
func main() {
	var n int
	var din float64

	fmt.Scan(&din)
	fmt.Println("Forma de pagamento")
	fmt.Scan(&n)
	 if n == 1 {
		din *= 0.9
		fmt.Printf("Pagamento a vista no valor de R$:%.2f\n", din)
	 } else if n == 2 {
		din *= 0.95
		fmt.Printf("Pagamento a vista no cartao no valor de R$:%.2f\n", din)
	 } else if n == 3 {
		din = din / 2
		fmt.Printf("Dividido em duas parcelas no valor de R$:%.2f cada, total: R$%.2f\n", din, din*2)
	 } else if n == 4 {
		dinTotal := din * 1.10
		parcela := dinTotal / 3
		fmt.Printf("Dividido em tres parcelas no valor de R$:%.2f cada, total: R$:%.2f\n", parcela, dinTotal)
	 } else {
		fmt.Println("Forma de pagamento inválida")
	 }
	}