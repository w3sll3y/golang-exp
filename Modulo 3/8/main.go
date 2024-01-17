package main

import (
	"errors"
	"fmt"
)

func main() {
	valor, err := sum(5, 10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("O valor eh %v", valor)
}

func sum(a, b int) (int, error) {
	if a+b >= 50 {
		return 0, errors.New("A soma eh maior que 50")
	}
	return a + b, nil
}
