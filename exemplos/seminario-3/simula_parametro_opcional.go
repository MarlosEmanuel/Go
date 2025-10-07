package main

import "fmt"

func saudacao(nome string, mensagens ...string) {
    if len(mensagens) > 0 {
        fmt.Printf("Olá, %s! %s\n", nome, mensagens[0])
    } else {
        fmt.Printf("Olá, %s!\n", nome)
    }
}

func main() {
    saudacao("Alice")
 	// Output: Olá, Alice!
    saudacao("Bob", "Bem-vindo de volta!")
	// Output: Olá, Bob! Bem-vindo de volta!

    // OBS: Go não suporta parâmetros opcionais (com valores default) nativamente, isso é apenas uma simulação, utilizando variadic functions.
    // Esta abordagem não é um método oficial para parâmetros opcionais em Go.
    // O objetivo aqui é apenas ilustrar o conceito.
}