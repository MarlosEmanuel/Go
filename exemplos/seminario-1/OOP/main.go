package main

import "fmt"

type Animal struct {
	Nome string
}

func (a Animal) Falar() {
	fmt.Println(a.Nome, "Faz um som.")
}

type Cachorro struct {
	Animal
}

func (c Cachorro) Latir() {
	fmt.Println(c.Nome, "Late!")
}

func main() {
	c := Cachorro{Animal{"Rex"}}
	c.Falar()
	c.Latir()
}