package main
import "fmt"

func main() {
 var kw, sal, custo_kw, custo_con, desconto float64

fmt.Println("Escreva o salário e o consumo de energia")
fmt.Scan(&sal, &kw)

custo_kw = sal * 0.7 / 100
custo_con = custo_kw * kw
desconto = custo_con * 0.9

fmt.Println("Custo por kw: ", custo_kw)
fmt.Println("Custo por consumo: ", custo_con)
fmt.Println("Custo com desconto: ", desconto)
}