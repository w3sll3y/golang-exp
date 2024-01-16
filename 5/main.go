package main

import "fmt"

const a = "Hellor World"

type ID int

var (
	b bool    = true
	c int     = 10
	d string  = "Wesley"
	e float64 = 1.1
	f ID      = 1
)

func main() {
	var meuArray [3]int
	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30

	// fmt.Println(len(meuArray) - 1)
	// fmt.Println(meuArray[0])
	// fmt.Println(meuArray[len(meuArray)-1])

	for i, valor := range meuArray {
		fmt.Printf("O valor de indice %d eh %d\n", i, valor)
	}
}
