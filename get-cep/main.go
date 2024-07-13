package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Address struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, url := range os.Args[1:] {
		req, err := http.Get(url)

		if err != nil {
			fmt.Fprintf(os.Stdout, "error in the request: %v\n", err)
		}

		defer req.Body.Close()

		res, err := io.ReadAll(req.Body)

		if err != nil {
			fmt.Fprintf(os.Stdout, "error in reading the response: %v\n", err)
		}

		var address Address
		if err := json.Unmarshal(res, &address); err != nil {
			fmt.Fprintf(os.Stdout, "error in unmarshal: %v\n", err)
		}

		file, err := os.Create("address.json")
		if err != nil {
			fmt.Fprintf(os.Stdout, "error in create file: %v\n", err)
		}

		defer file.Close()

		_, err = file.WriteString(string(res))

		if err != nil {
			fmt.Fprintf(os.Stdout, "error in writing the file: %v\n", err)
		}

		var address2 Address
		readFile, err := os.ReadFile("address.json")
		if err := json.Unmarshal(readFile, &address2); err != nil {
			panic(err)
		}

		fmt.Println(address2)
	}
}
