package main

import "fmt"

func main() {
	// ###########################
	// ## 1. INTEIROS COM SINAL ##
	// ###########################

	fmt.Println("--- INTEIROS COM SINAL (intN) ---")

	// Variáveis com tipos inteiros de tamanhos diferentes (8, 16, 32, 64 bits).
	// O int8 é ideal para números pequenos, economizando memória.

	var numeroPequeno int8 = -128
	fmt.Printf("int8 (mínimo): %d\n", numeroPequeno) // Imprime -128, o valor mínimo de um int8.

	var numeroMedio int16 = 32767
	fmt.Printf("int16 (máximo): %d\n", numeroMedio) // Imprime 32767, o valor máximo de um int16.

	// O tipo 'int' (sem tamanho especificado) geralmente é um alias para int32 ou int64,
	// dependendo da arquitetura do seu sistema (32 ou 64 bits).

	var numeroNormal int32 = 2147483647
	fmt.Printf("int32 (máximo): %d\n", numeroNormal)

	var numeroGrande int64 = 9223372036854775807
	fmt.Printf("int64 (máximo): %d\n\n", numeroGrande)

	// ##############################
	// ## 2. INTEIROS SEM SINAL (uintN) ##
	// ##############################

	fmt.Println("--- INTEIROS SEM SINAL (uintN) ---")
	// uintN armazena apenas valores positivos (incluindo zero), dobrando o limite superior.

	var byteSemSinal uint8 = 255
	fmt.Printf("uint8 (máximo) ou 'byte': %d\n", byteSemSinal) // byte é um alias para uint8.

	var numeroPositivo uint16 = 65535
	fmt.Printf("uint16 (máximo): %d\n", numeroPositivo)

	var numeroEnorme uint32 = 4294967295
	fmt.Printf("uint32 (máximo): %d\n", numeroEnorme)

	var numeroGigante uint64 = 18446744073709551615
	fmt.Printf("uint64 (máximo): %d\n\n", numeroGigante)

	// ###########################
	// ## 3. EXEMPLO PRÁTICO (RUNE) ##
	// ###########################

	// 'rune' é um alias para int32, usado para representar caracteres Unicode.
	// Isso é útil para lidar com emojis ou caracteres especiais.

	var meuRune rune = '😎' // O Go armazena o código Unicode do emoji.
	fmt.Printf("Tipo 'rune' (caractere): %c (Código Unicode: %d)\n", meuRune, meuRune)

	// O tipo 'byte' é um alias para uint8, e representa um byte de dados (útil em I/O).
	var meuByte byte = 'a'
	fmt.Printf("Tipo 'byte' (caractere): %c (Valor decimal: %d)\n", meuByte, meuByte)
}
