package main

import "fmt"

func main() {
	salarios := map[string]int{
		"Wesley": 1000,
		"Joao":   2000,
		"Maria":  3000,
	}

	// fmt.Println(salarios["Wesley"])
	// delete(salarios, "Wesley")
	// salarios["Wes"] = 5000
	// fmt.Println(salarios["Wes"])

	// sal := make(map[string]int)
	// sal1 := map[string]int{}
	// sal1["Wesley"] = 6000
	// sal["Victoria"] = 20000

	for nome, salario := range salarios {
		fmt.Printf("O salario de %s eh de %d\n", nome, salario)
	}

	for _, salario := range salarios {
		fmt.Printf("O salario eh de %d\n", salario)
	}
}
