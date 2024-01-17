package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	wesley := Cliente{
		Nome:  "Wesley",
		Idade: 25,
		Ativo: true,
	}
	// fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", wesley.Nome, wesley.Idade, wesley.Ativo)
	wesley.Ativo = false
	fmt.Println(wesley.Ativo)

}
