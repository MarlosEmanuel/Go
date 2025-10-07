package main

import "fmt"

type Numero interface {
	int | float64
}

func Processar[T Numero](valor T) {
	fmt.Printf("Processando valor de tipo %T: %v\n", valor, valor)
}

func main() {
	Processar(10)
	// Output: Processando valor de tipo int: 10
	Processar(3.14159)
	// Output: Processando valor de tipo float64: 3.14159
}
