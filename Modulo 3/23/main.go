package main

func main() {
	a := 10
	// b := 2
	// c := 3
	switch a {
	case 1:
		println("a")
	case 2:
		println("b")
	case 3:
		println("c")
	default:
		println("d")
	}
	// nao existe else if no go e nem if ternario
	// temos >, <, >=, <=, ==, !=, &&, ||
	// if a > b {
	// 	println(a)
	// } else {
	// 	println(b)
	// }
	// if a > b && c > a {
	// 	println("a > b && c > a")
	// }
}