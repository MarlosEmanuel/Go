package main

import "fmt"

func main() {

	// Declarar e Usar depois
	somador := func(x, y int) int {
		return x + y
	}
	resultado := somador(10, 5)
	fmt.Printf("1. Resultado da função anônima armazenada: %d\n", resultado)
	// Output: 1. Resultado da função anônima armazenada: 15

	
	// Declarar e Usar, os () chamam a função
	resultadoImediato := func(a, b int) int {
		return a * b
	}(20, 3)
	fmt.Printf("2. Resultado da função anônima executada imediatamente: %d\n", resultadoImediato)
	// Output: 2. Resultado da função anônima executada imediatamente: 60
}
