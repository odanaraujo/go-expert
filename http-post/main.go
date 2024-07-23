package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	c := http.Client{}
	//o body é um io de Reader. Ele ler um slice de byte
	jsonVar := bytes.NewBuffer([]byte(`{"name": "Dan"}`))
	resp, err := c.Post("https://google.com", "application/json", jsonVar)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	io.CopyBuffer(os.Stdout, resp.Body, nil)

}
