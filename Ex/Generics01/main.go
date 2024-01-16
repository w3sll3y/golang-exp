package main

type Number int

type MyNumber interface {
	// Number | int | float64
	~int | ~float64
}

func Subtrair[T MyNumber](a, b T) T {
	return a - b
}

func main() {
	m := map[string]int{"Wesley": 6, "Victoria": 10}
	m2 := map[string]float64{"Wesley": 6.2, "Victoria": 10.4}
	m3 := map[string]Number{"Wesley": 12, "Victoria": 10}
	a := Subtrair(m["Wesley"], m["Victoria"])
	b := Subtrair(m2["Wesley"], m2["Victoria"])
	c := Subtrair(m3["Wesley"], m3["Victoria"])
	println(Subtrair(500, 200))
	println(a)
	println(b)
	println(c)
}
