package main
import (
	f"fmt"
)
func main() {
	
	cdd:=0
	soma:= 0
	for i:=50;i<=70;i++ {
		if i%2==0{
			soma += i
			cdd++
		}
	}
	media := soma / cdd
	f.Printf("soma = %d\nMedia = %d", soma, media)
	
}