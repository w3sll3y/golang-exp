package main

import "fmt"

type Salario struct {
	Nome  string
	Valor float64
}

func (s *Salario) Saque(newValor float64) float64 {
	s.Valor -= newValor
	return s.Valor
}

func (s *Salario) Deposito(newValor float64) float64 {
	s.Valor += newValor
	return s.Valor
}

func main() {
	novaConta := Salario{Nome: "Wesley", Valor: 5000.00}
	fmt.Println(novaConta.Nome)
	novaConta.Nome = "Fernandes"
	fmt.Println(novaConta.Nome)
	novaConta.Saque(300.00)
	fmt.Println(novaConta.Valor)
	novaConta.Deposito(700.00)
	fmt.Println(novaConta.Valor)
}
