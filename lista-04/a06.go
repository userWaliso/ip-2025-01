package main
import "fmt"
func main() {
	var ints[]int
	for i:=100;i>=1;i-- {
		ints = append(ints, i)
	}
	fmt.Println(ints)
}