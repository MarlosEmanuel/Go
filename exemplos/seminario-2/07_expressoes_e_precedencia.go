package main

import "fmt"

func main() {
	fmt.Println("--- EXPRESSÕES E PRECEDÊNCIA DE OPERADORES ---")

	// Variáveis de demonstração
	var a int = 10
	var b int = 5
	var c int = 2
	var d int = 1

	fmt.Printf("Valores Iniciais: a=%d, b=%d, c=%d, d=%d\n\n", a, b, c, d)

	// ##############################
	// # Nível 5: *, /, %, <<, >>, &, &^ (Precedência Máxima) #
	// ##############################

	// Exemplo: Multiplicação tem precedência sobre a Adição.
	// Resultado: 10 + (5 * 2) = 20
	resultadoN5 := a + b*c
	fmt.Printf("Nível 5: a + b * c  (10 + 5 * 2) = %d\n", resultadoN5)

	// Exemplo Bitwise &^ (Bit clear): Limpa os bits definidos no segundo operando.
	// 10 (1010) &^ 2 (0010) = 8 (1000)
	resultadoBitwise := a &^ c
	fmt.Printf("Nível 5: Bit Clear (a &^ c) (10 &^ 2) = %d\n\n", resultadoBitwise)

	// ##############################
	// # Nível 4: +, -, |, ^ (Precedência Média-Alta) #
	// ##############################

	// Exemplo: Nível 5 (divisão) ocorre antes do Nível 4 (subtração).
	// Resultado: 10 - (5 / 2) = 10 - 2 = 8 (Divisão de inteiros)
	resultadoN4 := a - b/c
	fmt.Printf("Nível 4: a - b / c (10 - 5 / 2) = %d\n", resultadoN4)

	// Exemplo Bitwise | (OR):
	// 10 (1010) | 5 (0101) = 15 (1111)
	resultadoBitwise2 := a | b
	fmt.Printf("Nível 4: Bitwise OR (a | b) (10 | 5) = %d\n\n", resultadoBitwise2)

	// ##############################
	// # Nível 3: ==, !=, <, <=, >, >= (Comparações) #
	// ##############################

	// O nível 5 (multiplicação) e o nível 4 (adição) são avaliados primeiro.
	// Expressão: (10 + 5) > (2 * 1) -> 15 > 2 -> true
	resultadoN3 := a+b > c*d
	fmt.Printf("Nível 3: a + b > c * d (15 > 2) = %t\n", resultadoN3)

	// Exemplo de associatividade da esquerda para a direita (dentro do mesmo nível):
	// Nota: Em Go, é impossível encadear operadores de comparação dessa forma (ex: a < b < c),
	// mas a regra se aplica a operadores do mesmo nível, como Nível 5:
	// a * b / c será avaliado como (a * b) / c.

	// Exemplo Prático de Comparação:
	if a == 10 && b < 10 { // && é Nível 2
		fmt.Println("Nível 3: a é 10 e b é menor que 10.")
	} else {
		fmt.Println("Nível 3: A condição falhou.")
	}

	// ##############################
	// # Nível 2: && (AND Lógico) #
	// ##############################

	// Exemplo de Expressão Lógica Complexa (N3 > N2)
	// Expressão: (10 == 10) && (5 < 10)
	resultadoN2 := (a == 10) && (b < 10)
	fmt.Printf("\nNível 2: (a == 10) && (b < 10) = %t\n", resultadoN2)

	// ##############################
	// # Nível 1: || (OR Lógico) #
	// ##############################

	// || é o operador de menor precedência.
	// Expressão: (10 == 11) || (5 < 10)
	// O && (Nível 2) seria avaliado antes se estivesse na mesma linha.
	resultadoN1 := (a == 11) || (b < 10)
	fmt.Printf("Nível 1: (a == 11) || (b < 10) = %t\n", resultadoN1)

	// --------------------------------------------------------------------------

	fmt.Println("\n--- NÃO É PERMITIDA SOBRECARGA DE OPERADORES ---")
	// Exemplo de Aplicação (Go não permite a sobrecarga do operador '+')

	type Ponto struct {
		x, y int
	}

	// Tentar definir func (p1 Ponto) + (p2 Ponto) Ponto NÃO funciona em Go.
	// Em vez disso, usamos métodos e funções:

	p1 := Ponto{x: 1, y: 2}
	p2 := Ponto{x: 3, y: 4}

	// A adição de structs deve ser feita via função ou método explícito:
	somaPontos := func(pA, pB Ponto) Ponto {
		return Ponto{x: pA.x + pB.x, y: pA.y + pB.y}
	}

	pResultante := somaPontos(p1, p2)
	fmt.Printf("Soma de Structs (Ponto) via Função: {x: %d, y: %d}\n", pResultante.x, pResultante.y)
	fmt.Println("Isso demonstra a AUSÊNCIA de Sobrecarga de Operadores em Go.")
}
