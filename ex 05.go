package main

import "fmt"

func main() {
    var matriz [3][2]int

    for i := 0; i < 3; i++ {
        for j := 0; j < 2; j++ {
            fmt.Printf("Informe o valor para a posição [%d][%d]: ", i, j)
            fmt.Scan(&matriz[i][j])
        }
    }

    fmt.Println("Matriz resultante:")
    for i := 0; i < 3; i++ {
        for j := 0; j < 2; j++ {
            fmt.Printf("%d\t", matriz[i][j])
        }
        fmt.Println()
    }
}
