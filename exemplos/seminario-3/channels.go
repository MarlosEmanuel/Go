package main

import "fmt"

func main() {
	// 1. Cria um canal de inteiros (sincronizado)
	ch := make(chan int)

	// 2. Inicia uma Goroutine (o "worker")
	go func() {
		fmt.Println("Worker: Enviando dado para o canal...")
		ch <- 42 // Envia o valor 42
	}()

	// 3. Recebe o valor do canal (a Goroutine principal para aqui até receber)
	resultado := <-ch

	fmt.Println("Principal: Recebido:", resultado)
}
