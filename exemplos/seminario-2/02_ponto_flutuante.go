package main

import "fmt"

func main() {
	fmt.Println("--- PONTO FLUTUANTE (float32 e float64) ---")

	// float32 tem precisão de aproximadamente 7 casas decimais (uso menos comum).
	var pontoFlutuante32 float32 = 3.1415926535
	fmt.Printf("float32 (Pi): %.10f (Note a perda de precisão após 7 casas)\n", pontoFlutuante32)

	// float64 é o tipo de ponto flutuante padrão em Go (uso mais comum e mais preciso).
	var pontoFlutuante64 float64 = 2.718281828459045
	fmt.Printf("float64 (e): %.15f (Precisão muito maior)\n", pontoFlutuante64)

	// Exemplo de aplicação: calcular a área de um círculo
	var raio float64 = 5.0
	const pi float64 = 3.14159265359 // Constantes são ótimas para valores fixos.
	var area = pi * raio * raio

	fmt.Printf("\nCálculo de Área:\n")
	fmt.Printf("Raio: %.1f\n", raio)
	fmt.Printf("Área (pi * r * r): %.2f\n", area)
}
