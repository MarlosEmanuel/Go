package main

import (
	"errors"
	"fmt"
)

// verifica se um nome de usuário atende ao critério de tamanho.
func validarNomeUsuario(nome string) (bool, error) {
	if len(nome) < 5 {
		return false, errors.New("o nome de usuário deve ter pelo menos 5 caracteres")
	}
	return true, nil
}

func main() {
	nome := "joao"
	_, err := validarNomeUsuario(nome)
	if err != nil {
		fmt.Printf("Falha na validação para '%s': %s\n", nome, err)
		//Output: Falha na validação para 'joao': o nome de usuário deve ter pelo menos 5 caracteres
	}
}
