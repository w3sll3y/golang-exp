package main

import "fmt"

func main() {

	// func() {
	// rodar webserver
	// }()

	total := func() int {
		return sum(1, 2, 3, 4, 5, 6, 7) * 2
	}()

	fmt.Println(sum(1, 2, 4, 5, 6, 7, 8, 5, 3, 2, 4, 5, 6, 7, 8, 8, 3, 34, 53, 5325, 22624))
	fmt.Println("total===", total)
}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
