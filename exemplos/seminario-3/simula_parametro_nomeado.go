package main

import "fmt"

func exibirDetalhes(params map[string]interface{}) {
    fmt.Printf("Nome: %s, Idade: %d, Email: %s\n",
        params["Nome"].(string),
        params["Idade"].(int),
        params["Email"].(string),
    )
}

func main() {
    exibirDetalhes(map[string]interface{}{
        "Nome":  "Alice",
        "Idade": 30,
        "Email": "alice@example.com",
    })

    // OBS: Go não suporta parâmetros nomeados nativamente, isso é apenas uma simulação, utilizando um mapa.
    // Esta abordagem não é um método oficial para parâmetros nomeados em Go.
    // O objetivo aqui é apenas ilustrar o conceito.
}