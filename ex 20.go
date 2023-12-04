package main

import "fmt"

func main() {
    var n int
    fmt.Print("Informe o tamanho do array: ")
    fmt.Scan(&n)

    array := make([]int, n)

    fmt.Println("Informe os elementos do array:")
    for i := 0; i < n; i++ {
        fmt.Scan(&array[i])
    }

    ordenado := true

    for i := 1; i < n; i++ {
        if array[i-1] > array[i] {
            ordenado = false
            break
        }
    }

    if ordenado {
        fmt.Println("O array está ordenado em ordem crescente.")
    } else {
        fmt.Println("O array não está ordenado em ordem crescente.")
    }
}
