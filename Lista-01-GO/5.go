package main
import "fmt"

func main() {
	var conta, valor, consumo float64
	var tipo string

	fmt.Println(Escreva a conta, o consumo e o tipo)
	fmt.Scan(&conta, &consumo, &tipo)

	if (tipo == "r") {valor = consumo * 0.05 + 5
	fmt.Println("CONTA = ", conta,"\n", "VALOR A SER PAGO = ", valor)
}else {
	(tipo == "i") {valor = (consumo - 80) * 0.25 + 500
	fmt.Println("CONTA = ", conta,"\n", "VALOR A SER PAGO = ", valor)

}

}