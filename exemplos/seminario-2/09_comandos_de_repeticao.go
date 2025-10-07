package main

import "fmt"

func main() {
	fmt.Println("--- COMANDOS DE REPETIÇÃO (FOR) ---")

	// ##############################
	// # 1. FOR Clássico (C-like) #
	// ##############################

	fmt.Println("\n--- 1. FOR com Inicialização, Condição e Pós-Iteração ---")
	// Ideal para repetir um número FIXO de vezes.

	// SINTAXE: for inicializacao; condicao; pos-iteracao { }
	for i := 0; i < 5; i++ {
		// A variável 'i' tem escopo restrito a este loop 'for'.

		// Exemplo de 'continue': Pula para a próxima iteração se a condição for verdadeira.
		if i == 2 {
			fmt.Println("   CONTINUE: Pulando o número 2.")
			continue
		}

		fmt.Printf("   Iteração Clássica: %d\n", i)

		// Exemplo de 'break': Sai imediatamente do loop.
		if i == 4 {
			fmt.Println("   BREAK: Saindo do loop em 4.")
			break
		}
	}

	// ##############################
	// # 2. FOR Condicional (Simulando WHILE) #
	// ##############################

	fmt.Println("\n--- 2. FOR como WHILE (Apenas Condição) ---")
	// Ideal quando a contagem é controlada por uma variável externa ou uma condição complexa.

	tentativas := 0
	limite := 3

	// SINTAXE: for condicao { }
	for tentativas < limite {
		fmt.Printf("   Tentativa número: %d\n", tentativas+1)

		// Incremento da variável de controle DENTRO do corpo do loop.
		tentativas++
	}

	// ##############################
	// # 3. FOR Infinito #
	// ##############################

	fmt.Println("\n--- 3. FOR Infinito (Sempre Verdadeiro) ---")
	// Usado para servidores, listeners ou loops que requerem uma quebra explícita.

	// SINTAXE: for { }
	contador := 0
	for {
		if contador >= 2 {
			// É obrigatório usar 'break' para sair de um loop infinito.
			fmt.Println("   Break no loop infinito. Fim da execução.")
			break
		}
		fmt.Printf("   Loop Infinito, Contagem: %d\n", contador)
		contador++
	}

	// ##############################
	// # 4. FOR RANGE (Iteração sobre Coleções) #
	// ##############################

	fmt.Println("\n--- 4. FOR RANGE (Iteração sobre um Slice) ---")

	nomes := []string{"Marlos", "Eduardo", "Luiz"}

	// O for range retorna o índice e o valor (cópia) a cada iteração.
	for i, nome := range nomes {
		// O '_' é usado quando queremos IGNORAR um dos valores (ex: ignorar o índice).
		fmt.Printf("   Nome no Índice %d: %s\n", i, nome)
	}
}
