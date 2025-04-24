package main
import (
	f"fmt"
)
func main() {
	var x int
	f.Scan(&x)

	if x >= 3 {
		f.Print("0 1 ")
		t2:=0
		t1:=1
		for i:=1;i<x;i++ {
			t3:= t1+t2
			f.Printf("%d ", t3)
			t2 = t1
			t1 = t3
		}
	} else {
			f.Print("Número inválido")
			return
		}
}