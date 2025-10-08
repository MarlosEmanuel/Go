package main

import (
	"fmt"
	"time"
)

func main() {

	chMensagem := make(chan string)

	// O 'select' espera por eventos nos canais
	select {
	case msg := <-chMensagem:
		fmt.Println("Recebeu:", msg)
	case <-time.After(5 * time.Second):
		fmt.Println("Timeout!")
	}
}
