package main

import "fmt"

func main() {
	fmt.Println("--- ESTRUTURAS DE CONTROLE DE FLUXO ---")

	// ##############################
	// # 1. IF/ELSE (Com Inicializador de Variável) #
	// ##############################

	// Go permite declarar e inicializar uma variável (idade) antes da condição.
	// O escopo da variável 'idade' está restrito apenas ao bloco if/else.
	if idade := 20; idade >= 18 {
		fmt.Printf("IF: Idade verificada: %d. Você é maior de idade.\n", idade)
	} else {
		fmt.Printf("IF: Idade verificada: %d. Você é menor de idade.\n", idade)
	}

	// fmt.Println(idade) // Isso causaria um erro de compilação (idade está fora de escopo)

	// ##############################
	// # 2. FOR (O Único Loop em Go) #
	// ##############################

	fmt.Println("\n--- FOR: Tipo 1 (Clássico - C-like) ---")
	// Estrutura clássica: inicialização; condição; pós-iteração.
	for i := 0; i < 3; i++ {
		fmt.Printf("FOR Clássico: Iteração %d\n", i)
	}

	fmt.Println("\n--- FOR: Tipo 2 (Como WHILE) ---")
	// O for pode ser usado como um loop 'while', omitindo a inicialização e a pós-iteração.
	j := 0
	for j < 2 {
		fmt.Printf("FOR como WHILE: J = %d\n", j)
		j++
	}

	fmt.Println("\n--- FOR: Tipo 3 (For Range) ---")
	// O 'for range' é a forma idiomática de iterar sobre arrays, slices, maps ou canais.
	numeros := []int{10, 20, 30}

	// index (i) e value (v) são retornados a cada iteração.
	for i, v := range numeros {
		fmt.Printf("FOR RANGE: Índice %d, Valor %d\n", i, v)
	}

	// ##############################
	// # 3. SWITCH (Simples e Eficaz) #
	// ##############################

	fmt.Println("\n--- SWITCH: Avaliação Simples ---")
	// O switch em Go NÃO precisa de 'break', pois ele sai automaticamente após o 'case'.
	estado := "São Paulo"

	switch estado {
	case "Rio de Janeiro":
		fmt.Println("SWITCH: O estado é o Rio de Janeiro.")
	case "São Paulo":
		fmt.Println("SWITCH: O estado é São Paulo.")
	default:
		fmt.Println("SWITCH: Não encontrei este estado na lista.")
	}

	fmt.Println("\n--- SWITCH: Tipo 2 (Sem Expressão - Como Múltiplos IF/ELSE) ---")
	// Se nenhuma expressão for fornecida, ele avalia condições booleanas.
	pontuacao := 85

	switch {
	case pontuacao >= 90:
		fmt.Println("SWITCH (Sem Expressão): Excelente!")
	case pontuacao >= 80:
		fmt.Println("SWITCH (Sem Expressão): Bom trabalho!")
	default:
		fmt.Println("SWITCH (Sem Expressão): Precisa melhorar.")
	}
}
