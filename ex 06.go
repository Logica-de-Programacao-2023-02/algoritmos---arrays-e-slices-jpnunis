package main

import "fmt"

func main() {
    array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

    var valor int
    fmt.Print("Informe um valor: ")
    fmt.Scan(&valor)

    encontrado := false
    for _, v := range array {
        if v == valor {
            encontrado = true
            break
        }
    }

    fmt.Println("Resultado da busca:", encontrado)
}
