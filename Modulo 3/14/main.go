package main

import "fmt"

type Endereco struct {
	Logadouro string
	Numero    int
	Cidade    string
	Estado    string
}

type Pessoa interface {
	Desativar()
}

type Empresa struct {
	Nome string
}

func (e Empresa) Desativar() {}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	// Address Endereco
	Endereco
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado ", c.Nome)
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {
	wesley := Cliente{
		Nome:  "Wesley",
		Idade: 25,
		Ativo: true,
	}

	minhaEmpresa := Empresa{}
	// fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", wesley.Nome, wesley.Idade, wesley.Ativo)
	// wesley.Ativo = false
	// wesley.Desativar()
	// wesley.Endereco.Logadouro = "Altos"
	// fmt.Println(wesley.Ativo)
	Desativacao(minhaEmpresa)
	Desativacao(wesley)

}
