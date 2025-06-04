package main
import f "fmt"

func main() {
	var x, soma int
	f.Scan(&x)
	vetor := make([]int, x)
	array := make([]int, x)

	for i := 0; i < x; i++ {
		f.Scan(&vetor[i])
	}
	for i := 0; i < x; i++ {
		f.Scan(&array[i])
	}
	for i := 0; i < x; i++ {
		soma += vetor[i] * array[i]
	}
		
	f.Println(soma)


	
}