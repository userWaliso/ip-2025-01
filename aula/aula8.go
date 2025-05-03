package main
import f"fmt"


func fatorial(x int)int {
	if x < 1 {
		return 1
	} else {
		return x * fatorial(x-1)
	}

}
func main() {
	var fat int = 4
	result := fatorial(fat)
	f.Println(result)
}