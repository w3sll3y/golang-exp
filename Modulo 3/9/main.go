package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2, 4, 5, 6, 7, 8, 5, 3, 2, 4, 5, 6, 7, 8, 8, 3, 34, 53, 5325, 22624))
}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
