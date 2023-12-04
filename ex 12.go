package main

import "fmt"

func main() {
    array := [5]int{2, 6, 9, 12, 7}

    novoSlice := []int{}

    for _, valor := range array {
        if valor%3 == 0 {
            novoSlice = append(novoSlice, valor)
        }
    }

    fmt.Println("Novo Slice:", novoSlice)
}
