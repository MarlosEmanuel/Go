package main

import (
	"fmt"
	"time"
)

// Palavras-chave: type, struct (já visto, mas repetido para o exemplo de método)
type Operacao struct {
	Valor int
}

// Palavra-chave: func (Declaração de Função/Método)
// Palavra-chave: return (Retorna um valor)
func (o Operacao) dobrar() int {
	return o.Valor * 2
}

// -----------------------------------------------------

// Palavra-chave: go (Inicia uma Goroutine - concorrência)
func contagem(nome string) {
	// Palavra-chave: for (Loop de repetição)
	for i := 0; i < 3; i++ {
		// Palavra-chave: defer (Atrasar a execução de uma função)
		defer fmt.Printf("DEFER: %s finalizou o passo %d.\n", nome, i)
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("GOROUTINE: %s, passo %d\n", nome, i)
	}
}

func main() {
	fmt.Println("--- 1. ESTRUTURAS DE CONTROLE ---")

	// Palavra-chave: var (Declaração de Variável) - Já foi amplamente usada
	// Palavra-chave: const (Declaração de Constante)
	const limite = 5
	i := 0 // Declaração curta (Palavra-chave: var implícita)

	// Palavra-chave: if, else, else if (Estruturas Condicionais)
	if i < limite {
		fmt.Printf("A variável 'i' (%d) é menor que o limite (%d).\n", i, limite)
	} else if i == limite {
		fmt.Println("i é igual ao limite.")
	} else {
		fmt.Println("i é maior que o limite.")
	}

	// Palavra-chave: switch, case, default (Estruturas de Múltipla Escolha)
	dia := time.Now().Weekday()
	switch dia {
	// Palavra-chave: fallthrough (Quebra a regra de 'break' implícito - Raro)
	case time.Saturday, time.Sunday:
		fmt.Println("É final de semana!")
	default: // Palavra-chave: default
		fmt.Println("É dia útil.")
	}

	// -----------------------------------------------------

	fmt.Println("\n--- 2. FUNÇÕES, MÉTODOS E INTERFACES ---")

	// Palavra-chave: interface (Define um conjunto de métodos)
	type Dobravel interface {
		dobrar() int
	}

	op := Operacao{Valor: 7}
	var d Dobravel = op // op implementa implicitamente a interface Dobravel
	fmt.Printf("Método 'dobrar' (com receiver Operacao): %d\n", d.dobrar())

	// -----------------------------------------------------

	fmt.Println("\n--- 3. CONCORRÊNCIA E CANAIS ---")

	// Palavra-chave: chan (Declaração de Canal para comunicação entre Goroutines)
	c := make(chan string)

	// Palavra-chave: select (Seleciona uma operação de comunicação pronta)
	go func() {
		// Palavra-chave: continue (Pula para a próxima iteração do loop)
		for j := 0; j < 2; j++ {
			if j == 0 {
				continue
			}
			c <- "Mensagem enviada"
		}
	}()

	select {
	case msg := <-c: // Recebe do canal 'c'
		fmt.Printf("Recebido via 'select': %s\n", msg)
	// Palavra-chave: default (Se nenhum canal estiver pronto no momento)
	default:
		fmt.Println("Nenhuma mensagem pronta para ser recebida.")
	}

	// Inicia as Goroutines
	go contagem("Goroutine A")
	go contagem("Goroutine B")
	time.Sleep(500 * time.Millisecond) // Dá tempo para as Goroutines executarem
	// Atrasos de 'defer' serão impressos agora, na ordem LIFO (Last In, First Out)

	// -----------------------------------------------------

	fmt.Println("\n--- 4. CONTROLE AVANÇADO (GOTO) ---")
	// Palavra-chave: goto (Transferência incondicional de controle - EVITAR)

	fmt.Println("1. Antes do goto.")
	// Palavra-chave: goto (Salta para o label 'FIM')
	goto FIM

	// Palavra-chave: break (Sai imediatamente de um loop ou switch)
	for {
		fmt.Println("Este código NUNCA será executado.")
		break
	}

	// Palavra-chave: case (Em switch)

FIM: // Palavra-chave: label (Usado com 'goto' e 'break/continue' aninhados)
	fmt.Println("2. Chegou no label FIM (após o goto).")

	// Palavras-chave restantes (já abordadas ou implícitas):
	// var, func, package, import, bool, string, int, float, byte, rune,
	// range, map, struct, make, new, append, len, cap, nil, error (interface)
}

// Palavra-chave: package (Deve ser 'main' para ser executável)
// Palavra-chave: import (Importa pacotes)
