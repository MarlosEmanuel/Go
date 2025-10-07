package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Atoi converte uma string para int e retorna um erro se não for possível.
	i, err := strconv.Atoi("3a")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
		return
	}
	fmt.Println("Número convertido:", i)
	// Output: Ocorreu um erro: strconv.Atoi: parsing "3a": invalid syntax
}
