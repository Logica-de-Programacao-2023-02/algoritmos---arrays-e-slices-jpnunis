package main

import "fmt"

func main() {
    array := [10]float64{2.3, 6.7, 8.1, 3.2, 9.5, 1.0, 7.8, 4.2, 5.6, 10.1}

    novoSlice := []float64{}

    for _, valor := range array {
        if valor > 5 {
            novoSlice = append(novoSlice, valor)
        }
    }

    fmt.Println("Novo Slice:", novoSlice)
}
