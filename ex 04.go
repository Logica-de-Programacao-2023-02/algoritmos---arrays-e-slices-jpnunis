package main

import "fmt"

func main() {
    var tamanho int
    fmt.Print("Informe o tamanho do Slice: ")
    fmt.Scan(&tamanho)

    slice := make([]int, tamanho)

    for i := 0; i < tamanho; i++ {
        fmt.Printf("Informe o valor para o elemento %d: ", i+1)
        fmt.Scan(&slice[i])
    }

    soma := 0
    for _, valor := range slice {
        soma += valor
    }

    fmt.Println("Slice:", slice)
    fmt.Println("Soma dos valores:", soma)
}
