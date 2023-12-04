package main

import (
	"fmt"
	"math"
)

func main() {
	slice := []int{3, 1, 8, 4, 6, 2, 10, 7, 9, 5}

	minimo, maximo := math.MaxInt, math.MinInt

	for _, valor := range slice {
		if valor < minimo {
			minimo = valor
		}
		if valor > maximo {
			maximo = valor
		}
	}

	fmt.Printf("Valor mínimo: %d\n", minimo)
	fmt.Printf("Valor máximo: %d\n", maximo)
}
