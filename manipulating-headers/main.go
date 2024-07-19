package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const url = "https://viacep.com.br/ws/%s/json/"

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
	http.HandleFunc("/", BuscaCepHandler)
	http.ListenAndServe(":8080", nil)
}

func BuscaCepHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	cepParam := r.URL.Query().Get("cep")

	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("zip code invalid"))
		return
	}

	adress, err := getCep(cepParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	//toda essa linha comentada
	/*
		result, err := json.Marshal(adress)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(result)
	*/

	w.WriteHeader(http.StatusOK)
	//é exatamente essa linha aqui
	if err := json.NewEncoder(w).Encode(adress); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
}

func getCep(zipCode string) (*Address, error) {
	req, err := http.Get(fmt.Sprintf(url, zipCode))

	if err != nil {
		return nil, errors.New("error in the request")
	}

	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)

	if err != nil {
		return nil, errors.New("error in reading the response")
	}

	var address Address
	if err := json.Unmarshal(res, &address); err != nil {
		return nil, errors.New("error in unmarshal")
	}

	return &address, nil
}
