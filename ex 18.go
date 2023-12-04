package main

import "fmt"

func main() {
    var n int
    fmt.Print("Informe um número inteiro positivo: ")
    fmt.Scan(&n)

    if n <= 0 {
        fmt.Println("Por favor, informe um número inteiro positivo.")
        return
    }

    primos := make([]int, 0, n)
    numero := 2

    for len(primos) < n {
        if isPrimo(numero) {
            primos = append(primos, numero)
        }
        numero++
    }

    fmt.Println("Array com os", n, "primeiros números primos:", primos)
}

func isPrimo(num int) bool {
    if num < 2 {
        return false
    }

    for i := 2; i*i <= num; i++ {
        if num%i == 0 {
            return false
        }
    }

    return true
}
