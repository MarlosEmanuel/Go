package main

import "fmt"

func IncrementarValor(numero int) {
	numero = numero + 10
	fmt.Printf("Valor dentro da função: %d\n", numero) // Output: Valor dentro da função: 15
}

func main() {

	variavelOriginal := 5

	fmt.Printf("Valor Original (antes da função): %d\n", variavelOriginal) // Output: Valor Original (antes da função): 5

	IncrementarValor(variavelOriginal)

	fmt.Printf("Valor Original (depois da função): %d\n", variavelOriginal) // Output: Valor Original (depois da função): 5
}
