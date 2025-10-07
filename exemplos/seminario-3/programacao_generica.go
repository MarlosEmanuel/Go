package main

import "fmt"

// T é um parâmetro de tipo que aceita int ou float64
func Soma[T int | float64](a, b T) T {
	return a + b
}

func main() {
	fmt.Println("Soma de inteiros:", Soma(5, 10))
	// Output: Soma de inteiros: 15

	fmt.Println("Soma de floats:", Soma(3.14, 2.71))
	// Output: Soma de floats: 5.85

}
