package main

func main() {
	flecha := "30Km/h"
	println(flecha)

	var ponteiroFlecha *string = &flecha
	*ponteiroFlecha = "40Km/h"
	println(flecha)
	novoPontFlecha := ponteiroFlecha
	*novoPontFlecha = "60Km/h"
	println(flecha)

	// & -> Passa endereco para nova varivavel ou mostra endereco de uma varival // * -> Valor do endereco de uma variavel que ja pegou o endereco de outra variavel
	println(*ponteiroFlecha)
}
