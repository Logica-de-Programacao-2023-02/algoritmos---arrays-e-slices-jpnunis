package main

import "fmt"

func main() {
    var n int
    fmt.Print("Informe o tamanho dos arrays: ")
    fmt.Scan(&n)

    array1 := make([]int, n)
    array2 := make([]int, n)

    fmt.Println("Informe os elementos do primeiro array:")
    for i := 0; i < n; i++ {
        fmt.Scan(&array1[i])
    }

    fmt.Println("Informe os elementos do segundo array:")
    for i := 0; i < n; i++ {
        fmt.Scan(&array2[i])
    }

    somaArray := make([]int, n)

    for i := 0; i < n; i++ {
        somaArray[i] = array1[i] + array2[i]
    }

    fmt.Println("Terceiro array (soma):", somaArray)
}
