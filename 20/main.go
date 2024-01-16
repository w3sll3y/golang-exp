package main

type MyNumber int

type Number interface {
	~int | ~float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

// comparable so eh utilizada a nivel de comparacao, por parecer com any, mas como string pode ser comparada
// e nao verificada se eh > que outra, o comparable nao eh usada para verificar > ou <
func Compara[T comparable](a, b T) bool {
	if a == b {
		return true
	}
	return false
}

// func SomaInteiro(m map[string]int) int {
// 	var soma int
// 	for _, v := range m {
// 		soma += v
// 	}
// 	return soma
// }

// func SomaFloat(m map[string]float64) float64 {
// 	var soma float64
// 	for _, v := range m {
// 		soma += v
// 	}
// 	return soma
// }

func main() {
	m := map[string]int{"Wesley": 1000, "Joao": 2000, "Maria": 3000}
	m2 := map[string]float64{"Wesley": 100.20, "Joao": 2000.3, "Maria": 300.0}
	// println(SomaInteiro(m))
	// println(SomaFloat(m2))
	m3 := map[string]MyNumber{"Wesley": 1000, "Joao": 2000, "Maria": 3000}

	println(Soma(m))
	println(Soma(m2))
	println(Soma(m3))

	println(Compara(10, 10.0))
}
