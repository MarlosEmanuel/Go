package main

import (
	"errors"
	"fmt"
)

func validarIdade(idade int) error {
	if idade < 18 {
		// Criando um novo erro
		return errors.New("usuário é menor de idade")
	}
	return nil
}

func cadastrarUsuario(idade int) error {
	err := validarIdade(idade)
	if err != nil {
		// Não tratei aqui, apenas deleguei (retornei) o erro
		return err
	}
	// ... continua o cadastro
	return nil
}

func main() {
	err := cadastrarUsuario(17)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
}
