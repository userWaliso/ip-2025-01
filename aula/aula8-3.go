package main
import f"fmt"


func invert(x []int, n int) {
	if n <= 1 {
		return 
	} 

	x[0], x[n-1] = x[n-1], x[0]

	invert(x[1:n-1], n-2)
}

func main() {
	var vetor = [5]int{1,2,3,4,5}
	invert(vetor[:], len(vetor))
	f.Printf("%v\n", vetor)
}