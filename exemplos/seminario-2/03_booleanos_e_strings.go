package main

import "fmt"

func main() {
	fmt.Println("--- TIPOS PRIMITIVOS BÁSICOS ---")

	// ###################
	// ## 1. BOOLEANOS (bool) ##
	// ###################

	var verdadeiro bool = true
	var falso bool = false

	fmt.Printf("Variável Booleana 1: %t\n", verdadeiro)
	fmt.Printf("Variável Booleana 2: %t\n", falso)

	// Exemplo de aplicação: controle de fluxo (muito mais aplicável do que apenas declarar)
	var usuarioLogado bool = true

	if usuarioLogado == verdadeiro { // ou apenas if usuarioLogado {
		fmt.Println("Usuário está logado. Acesso concedido.")
	} else {
		fmt.Println("Usuário deslogado. Redirecionando para login.")
	}

	// ###################
	// ## 2. STRINGS (string) ##
	// ###################

	var saudacao string = "Olá, Golang!"
	var linguagem string = "Go"

	// As strings em Go são imutáveis e representam sequências de bytes (UTF-8).
	fmt.Printf("\nString: %s\n", saudacao)

	// Concatenando strings:
	var mensagemCompleta = saudacao + " Este é um seminário sobre " + linguagem + "."
	fmt.Printf("Concatenação: %s\n", mensagemCompleta)

	// Obtendo o tamanho da string (número de bytes):
	fmt.Printf("Tamanho da string 'Olá, Golang!' (em bytes): %d\n", len(saudacao))
}
