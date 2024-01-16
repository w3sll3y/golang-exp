package main

import "fmt"

type Rua struct {
	Endereco string
	Numero   int
}

type Porta struct {
	Cor     string
	Tamanho int
	Fechada bool
	Rua
}

func (p Porta) Fechar() bool {
	p.Fechada = false
	fmt.Printf("Minha porta estar PPPP %v\n", p.Fechada)
}

func main() {
	minhaPorta := Porta{
		Cor:     "Azul",
		Tamanho: 10,
		Fechada: true,
	}
	minhaPorta.Rua.Endereco = "R 01"
	minhaPorta.Rua.Numero = 12
	fmt.Printf("Minha porta estar %v\n", minhaPorta.Fechada)
	minhaPorta.Fechar()
	fmt.Println(minhaPorta.Fechada)
}
