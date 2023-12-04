package main

import "fmt"

func main() {
    slice := make([]int, 0, 5)

    var numero int
    fmt.Print("Informe um número inteiro: ")
    fmt.Scan(&numero)

    presente := false
    for _, v := range slice {
        if v == numero {
            presente = true
            break
        }
    }

    if !presente {
        slice = append(slice, numero)
    }

    fmt.Println("Slice resultante:", slice)
}
