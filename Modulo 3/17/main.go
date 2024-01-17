package main

type Conta struct {
	saldo int
}

func (c *Conta) Simular(valor int) int {
	c.saldo += valor
	println(c.saldo)
	return c.saldo
}

func newConta() *Conta {
	return &Conta{saldo: 0}
}

func main() {
	conta := Conta{saldo: 100}
	conta.Simular(200)
	println(conta.saldo)
}
