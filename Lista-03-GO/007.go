package main

import (
	f "fmt"
)

func main() {
	var (
		x, y int
		media, mediaPar, per, num, par, soma, impar float64
	)
	numeros := []int{}
	pares := []int{}
	f.Scan(&x)
	for i := 0; i < x; i++ {
		f.Scan(&y)
		num++
		soma += float64(y)
		if y%2 == 0 {
			pares = append(pares, y)
			par++
		} else {
			impar++
		}
		numeros = append(numeros, y)
	}
	somaPar := 0.0
	for _, v := range pares {
		somaPar += float64(v)
	}

	min := numeros[0]
	max := numeros[0]

	for _, n := range numeros {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	mediaPar = somaPar / par
	media = soma / num
	per = (impar * 100) / num

	f.Printf("a soma dos números digitados: %d\n", int(soma))
	f.Printf("a quantidade de números digitados: %d\n", x)
	f.Printf("a média dos números digitados: %.1f\n", media)
	f.Printf("o maior número digitado: %d\n", max)
	f.Printf("o menor número digitado: %d\n", min)
	f.Printf("a média dos números pares: %.1f\n", mediaPar)
	f.Printf("a percentagem dos números ímpares entre todos os números digitados: %.2f%%\n", per)
}
