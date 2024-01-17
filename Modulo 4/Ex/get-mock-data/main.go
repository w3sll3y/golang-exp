package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Pessoas struct {
	CreatedAt string `json:"createdAt"`
	Name      string `json:"name"`
	Avatar    string `json:"avatar"`
	ID        string `json:"id"`
}

func main() {
	for _, id := range os.Args[1:] {
		req, err := http.Get("https://65a82e3f94c2c5762da86d28.mockapi.io/go/v2/data/" + id)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao buscar dados: %s\n", err)
		}
		defer req.Body.Close()
		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao passar dados %s\n", err)
		}
		fmt.Println("RESSS===", string(res))
		var data Pessoas
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao passar para data %s\n", err)
		}
		file, err := os.Create("pessoa.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao cadastrar arquivo, %s\n", err)
		}
		defer file.Close()
		_, err = file.WriteString(fmt.Sprintf("Nome: %s, ID: %s, Avatar: %s, Data de criacao: %.2f\n", data.Name, data.ID, data.Avatar, data.CreatedAt))
		fmt.Println("Arquivo criado com sucesso")
		fmt.Println("Nome", data.Name)
	}
}
