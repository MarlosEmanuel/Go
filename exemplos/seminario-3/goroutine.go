package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("Rodando em paralelo")
	fmt.Println("Rodando normalmente")
	time.Sleep(time.Second)
}
