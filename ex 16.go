package main

import "fmt"

func main() {
    array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

    novoSlice := []int{}

    for _, valor := range array {
        if valor%2 == 0 {
            novoSlice = append(novoSlice, valor)
        }
    }

    fmt.Println("Novo Slice:", novoSlice)
}
