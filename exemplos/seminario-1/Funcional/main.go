package main

import "fmt"

func contador () func () int {
	x := 0 
	return func () int {
		x++
		return x
	}
}

func main () {
	c := contador ()
	fmt.Println(c())
	fmt.Println(c())
}