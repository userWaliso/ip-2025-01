package main
import f"fmt"
func main() {
	var impar[]int
	for i:=1;i<200;i++ {
		if i%2 != 0 {
			impar = append(impar, i)
		}
	}
	f.Println(impar)
}