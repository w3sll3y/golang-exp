package main

import "fmt"

func main() {
	var minhaVar interface{} = "Wesley Fernandes"
	println(minhaVar.(string))
	res, ok := minhaVar.(int)
	fmt.Printf("o valor de res eh %v e o resultado de ok eh %v", res, ok)

	// res2 := minhaVar.(int)
	// println(res2)
}
