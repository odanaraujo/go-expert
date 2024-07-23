package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	//como deixar algo mais performatico?
	//Estabelecer os limites das chamadas externas que o sistema vai utilizar
	//uma api externa demora 10s para executar.
	//É preciso sempre ter em mente esse tempo limite que a nossa aplicação fará a execução

	c := http.Client{Timeout: 1 * time.Second}
	resp, err := c.Get("https://google.com")
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
