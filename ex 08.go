package main

import "fmt"

func main() {
    slice := []string{"apple", "banana", "orange", "apple", "grape", "kiwi", "apple", "mango"}

    var valor string
    fmt.Print("Informe um valor: ")
    fmt.Scan(&valor)

    novoSlice := make([]string, 0, len(slice))

    for _, v := range slice {
        if v != valor {
            novoSlice = append(novoSlice, v)
        }
    }

    fmt.Println("Slice resultante:", novoSlice)
}
