package main
import (
	f"fmt"
)
func main() {
	var (
		y,impares int
		vet = []int{}
	)
	for i:=0;i<10;i++ {
		f.Scan(&y)
		vet = append(vet, y)
		if y%2 != 0 {
			impares++
		}
	}
	somaPar := 0
	
	var (
		par []int
		impar []int
	)
	for _,v := range vet {
		if v%2 == 0 {
			somaPar += v
			par = append(par, v)
			} 
		}
		f.Printf("numeros pares: %d\n", par)
	for _,v := range vet {
		if v%2 != 0 {
			impar = append(impar, v)
		}
	}
	f.Printf("numeros impares: %d\n", impar)
	f.Printf("soma dos números pares: %d\nquantidade de números ímpares: %d\n",somaPar, impares)
}