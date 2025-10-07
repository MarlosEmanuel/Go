package main

import (
	"fmt"
	"time"
)

// Variável Global: Inicia com letra maiúscula, tornando-a PÚBLICA (exportada)
var NomeDoSistema string = "SistemaGoSeminario"

// Variável Privada: Inicia com letra minúscula, visível apenas neste pacote.
var contadorGlobal int = 0

// ##############################
// # Demonstração de Case Sensitive, Minúsculas e Maiúsculas #
// ##############################

// Função Pública (Exportada): Inicia com letra maiúscula.
func ExibirRegras() {

	// 1. Case Sensitive
	var variavel int = 10
	var Variavel int = 20 // Go as trata como duas variáveis COMPLETAMENTE diferentes.

	fmt.Println("--- 1. CASE SENSITIVE ---")
	fmt.Printf("variavel: %d, Variavel: %d\n", variavel, Variavel)

	// 2. Palavras-chave
	// Você NÃO pode usar palavras-chave reservadas como nomes de variáveis.
	// var break int = 5 // Isso causaria um erro de compilação.

	// 3. Vinculação (Binding) e Tempo de Vida
	// Variável local, vinculada e com tempo de vida restrito a esta função.
	nomeLocal := "Eduardo"
	fmt.Printf("\n--- 2. VINDAÇÃO E ESCOPO LOCAL ---\n")
	fmt.Printf("Variável 'nomeLocal' (dentro da função): %s\n", nomeLocal)

	// Incrementa a variável privada do pacote
	contadorGlobal++
	fmt.Printf("Contador global após incremento: %d\n\n", contadorGlobal)

}

// Função Privada (Não Exportada): Inicia com letra minúscula.
func calcularMedia(nota1, nota2 float64) float64 {

	// Demonstração de Escopo de Bloco
	// Variável 'soma' existe apenas dentro da estrutura 'if'.
	if nota1 > 0 {
		soma := nota1 + nota2 // Vinculação de 'soma' a este bloco.
		return soma / 2.0
	}
	// fmt.Println(soma) // Se descomentar, causa um erro: 'soma' é indefinida aqui.
	return 0
}

// ##############################
// # Demonstração de Escopo e Vinculação #
// ##############################

func main() {

	fmt.Println("--- 3. CARACTERE INICIAL (VISIBILIDADE) ---")

	// Acessando a variável pública de outro pacote (simulando a importação)
	// Go usa o Caractere Inicial Maiúsculo para indicar Exportação/Público.
	fmt.Printf("Nome do Sistema (Global, Público): %s\n", NomeDoSistema)

	// Tentativa de acesso à variável privada:
	// fmt.Println(contadorGlobal) // É acessível pois estamos no MESMO pacote (main).
	// Se fosse importada, não seria acessível!

	// Executando a demonstração de regras
	ExibirRegras()

	// -----------------------------------------------------

	fmt.Println("--- 4. TEMPO DE VIDA E GARBAGE COLLECTOR ---")

	// Usando o tipo 'time.Time' importado
	agora := time.Now()

	// A função 'time.Now()' retorna um objeto 'Time'.
	// Quando 'main' terminar, o Go's Garbage Collector irá liberar
	// automaticamente a memória alocada para 'agora' e 'mediaCalculada'.
	mediaCalculada := calcularMedia(9.5, 8.0)

	fmt.Printf("Horário de Início (Variável 'agora'): %v\n", agora.Format(time.Kitchen))
	fmt.Printf("Média Calculada: %.2f\n", mediaCalculada)

	fmt.Println("\nFim do programa. A memória das variáveis locais será limpa.")
}
