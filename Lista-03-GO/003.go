package main
import (
	f"fmt"
)
func main() {
	var carlos, joao float64
	f.Scan(&carlos)
	joao = carlos / 3
	cdd := 0
	for i:=0;joao <= carlos; i++ {
		carlos *= 1.02
		joao *= 1.05
		cdd++
	}
	f.Printf("sal do joao: %.2f sal do carlos: %.2f meses: %d\n",joao, carlos, cdd)
}