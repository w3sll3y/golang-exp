package main

import "fmt"

func main() {

	// Variavel -> ponteiro que tem um endereco na memoria -> valor
	a := 10
	var ponteiro *int = &a
	*ponteiro = 20
	b := &a
	*b = 30
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(b)
	fmt.Println(*b)
	fmt.Println(&b)
	fmt.Println(ponteiro)
	fmt.Println(*ponteiro)
	fmt.Println(&ponteiro)
}
