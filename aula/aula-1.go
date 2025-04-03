package main
import "fmt"

func main() {
	var aluno,media,soma,i float64

	fmt.Println("Digite a quantidade de alunos: ")
	fmt.Scan(&aluno)

	for i = 0; i < aluno; i++ {
		fmt.Printf("Digite a nota do aluno %.0f°: ", i+1)
		fmt.Scan(&media)
		soma += media
	}
	mediaglobal := soma/aluno
	fmt.Printf("A média global é %.1f\n: ",mediaglobal)
}