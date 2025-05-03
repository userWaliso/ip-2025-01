package main
import f"fmt"


func potencia(x int, n int)int {
	if n < 1 {
		return 1
	} else {
		return x * potencia(x, n-1)
	}

}
func main() {
	var x, n int
	f.Scan(&x, &n)
	result := potencia(x, n)
	f.Println(result)
}