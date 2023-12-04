package main

import "fmt"

func main() {
	numeros := [3]int{10, 20, 30}
	soma := 0

	for _, num := range numeros {
		soma += num
	}

	fmt.Printf("Soma dos valores no array: %d\n", soma)
}
