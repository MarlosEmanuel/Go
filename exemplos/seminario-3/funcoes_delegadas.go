package main

import "fmt"

type Operacao func(a, b int) int

func Somar(x, y int) int {
	return x + y
}

func Multiplicar(x, y int) int {
	return x * y
}

func ExecutarOperacao(calculo Operacao, a, b int) int {
	resultado := calculo(a, b)
	return resultado
}

func main() {
	valorA := 10
	valorB := 5

	resultadoSoma := ExecutarOperacao(Somar, valorA, valorB)
	fmt.Printf("Resultado da Soma: %d\n", resultadoSoma)
	// Output: Resultado da Soma: 15

	resultadoMultiplicacao := ExecutarOperacao(Multiplicar, valorA, valorB)
	fmt.Printf("Resultado da Multiplicação: %d\n", resultadoMultiplicacao)
	// Output: Resultado da Multiplicação: 50

	minhaFuncao := Somar
	resultadoVariavel := minhaFuncao(20, 3)
	fmt.Printf("Resultado da Variável de Função: %d\n", resultadoVariavel) 
	// Output: Resultado da Variável de Função: 23
}
