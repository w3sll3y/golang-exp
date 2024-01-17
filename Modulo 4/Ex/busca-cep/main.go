package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, cep := range os.Args[1:] {
		req, err := http.Get("https://viacep.com.br/ws/" + cep + "/json")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao buscar cep %s", cep)
		}
		defer req.Body.Close()
		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ao ler a responsta %v\n", err)
		}
		var dados ViaCEP
		err = json.Unmarshal(res, &dados)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao dar o parse dos dados %s\n", err)
		}
		file, err := os.Create("cidade.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao criar arquivo %s", err)
		}
		defer file.Close()
		_, err = file.WriteString(fmt.Sprintf("LOCALIDADE: %s, CEP: %s, BAIRRO: %s", dados.Localidade, dados.Cep, dados.Bairro))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao gravar no arquivo %s", err)
		}
		fmt.Println("Arquivo gravado com sucesso")
		fmt.Println("BAIRRO", dados.Cep)
	}
}
