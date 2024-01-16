package main

import "fmt"

type Endereco struct {
	Logadouro string
	Numero    int
	Cidade    string
	Estado    string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	// Address Endereco
	Endereco
}

func main() {
	wesley := Cliente{
		Nome:  "Wesley",
		Idade: 25,
		Ativo: true,
	}
	// fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", wesley.Nome, wesley.Idade, wesley.Ativo)
	wesley.Ativo = false
	wesley.Endereco.Logadouro = "Altos"
	fmt.Println(wesley.Ativo)

}
