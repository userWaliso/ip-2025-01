package main
import f"fmt"


func soma(x []int, n int)int {
	if n <= 0 {
		return 0
	} else {
		return x[n-1] + soma(x, n-1)
	}

}
func main() {
	var vetor = [5]int{2,3,4,5,6}
	result := soma(vetor[:], len(vetor))
	f.Println(result)
}