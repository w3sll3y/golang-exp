package main

import (
	"fmt"
	"golang-exp/21/matematica"

	"github.com/google/uuid"
)

func main() {
	s := matematica.Soma(10, 20)
	carro := matematica.Carro{Marca: "Fiat"}
	fmt.Printf("O resultado: %v\n", s)
	fmt.Println(matematica.A)
	fmt.Println(carro.Marca)
	fmt.Println(carro.Andar())
	fmt.Println(uuid.New())
}
