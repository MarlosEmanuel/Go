package main

import "fmt"

func main() {
	fmt.Println("--- 1. ARRAYS (Tamanho Fixo) ---")
	// Arrays são estáticos e raramente usados diretamente em Go, exceto para casos de performance fixa.

	var arrNumeros [5]int // Declara um Array de 5 inteiros, inicializado com zeros.
	arrNumeros[0] = 100   // Atribuição em um índice específico.
	fmt.Printf("Array de 5 inteiros (arrNumeros): %v\n", arrNumeros)

	arrCores := [3]string{"Vermelho", "Verde", "Azul"} // Declaração e inicialização em uma linha.
	fmt.Printf("Array de 3 strings (arrCores): %v\n", arrCores)
	fmt.Printf("Tamanho do Array: %d\n\n", len(arrCores))

	// -----------------------------------------------------

	fmt.Println("--- 2. SLICES (Dinâmicos e Flexíveis) ---")
	// Slices são a forma dinâmica e preferencial de trabalhar com listas em Go.

	sliceNumeros := []int{1, 2, 3} // Cria um Slice inicial.
	fmt.Printf("Slice Inicial: %v\n", sliceNumeros)

	// O 'append' cria um novo slice (se necessário) e adiciona o elemento.
	sliceNumeros = append(sliceNumeros, 4, 5)
	fmt.Printf("Slice após append (4, 5): %v\n", sliceNumeros)
	fmt.Printf("Tamanho do Slice: %d\n\n", len(sliceNumeros))

	// -----------------------------------------------------

	fmt.Println("--- 3. STRUCTS (Estruturas de Dados Customizadas) ---")

	// Declaração da Struct (melhor fazer fora da função, mas funciona aqui para o exemplo)
	type Pessoa struct {
		Nome  string
		Idade int
		Cargo string
	}

	// Cria uma nova instância da Struct
	aluno := Pessoa{
		Nome:  "Marlos Fontes",
		Idade: 25,
		Cargo: "Estudante", // Exemplo criativo: cargo no projeto de seminário
	}

	fmt.Printf("Struct Aluno: %+v\n", aluno) // %+v imprime o nome do campo e o valor.
	fmt.Printf("%s tem %d anos e é %s.\n\n", aluno.Nome, aluno.Idade, aluno.Cargo)

	// -----------------------------------------------------

	fmt.Println("--- 4. MAPS (Dicionários/Arrays Associativos) ---")
	// Maps armazenam pares chave-valor.

	m := make(map[string]int) // Cria um Map com chaves string e valores inteiros.
	m["seminario"] = 42
	m["disciplina"] = 101

	fmt.Printf("Valor da chave 'seminario': %d\n", m["seminario"])

	// Exemplo de verificação de existência (muito comum em Go):
	valor, existe := m["disciplina"]
	if existe {
		fmt.Printf("A chave 'disciplina' existe e tem o valor: %d\n", valor)
	}

	delete(m, "disciplina") // Remoção de um elemento.
	fmt.Printf("Map após remoção: %v\n\n", m)

	// -----------------------------------------------------

	fmt.Println("--- 5. PONTEIROS E REFERÊNCIAS ---")
	// Ponteiros permitem que as funções modifiquem o valor real de uma variável.

	var x int = 10
	var p *int = &x // '&' obtém o endereço de memória de 'x' (referência).

	fmt.Printf("Valor inicial de x: %d\n", x)
	fmt.Printf("Endereço de x (armazenado em p): %p\n", p)

	*p = 20 // '*' desreferencia o ponteiro 'p' e altera o valor no endereço de memória.

	fmt.Printf("Valor de x após desreferenciação (*p = 20): %d\n", x) // x agora é 20.
}
