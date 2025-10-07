package main

import "fmt"

func somar(a int, b int) int {
    return a + b
}

func main() {
    resultado := somar(5, 3)
    fmt.Println("Soma:", resultado)
}