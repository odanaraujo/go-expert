package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Conta struct {
	Number  int
	Balance int
}

func main() {
	c := Conta{
		Number:  1234,
		Balance: 100,
	}

	res, err := json.Marshal(c) //retorno em bytes
	if err != nil {
		panic(err)
	}

	file, err := os.Create("file.txt")

	defer file.Close()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(res))

	// podemos trabalhar com encoded, onde ele retorna o valor e já salva no arquivo, por exemplo, ou na tela do terminal )
	if err := json.NewEncoder(file).Encode(c); err != nil {
		panic(err)
	}

	openFile, err := os.Open("file.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(openFile)
	buffer := make([]byte, 10)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Printf("%s", string(buffer[:n]))
	}
	defer openFile.Close()

	// buscando um json e realizando o dencoder
	open, _ := os.Open("conta.json")
	open2, _ := os.ReadFile("conta.json")

	var cc Conta
	var cc2 Conta

	if err := json.NewDecoder(open).Decode(&cc); err != nil {
		panic(err)
	}

	fmt.Println(cc)

	//outra opção é usando unmarshal
	if err := json.Unmarshal(open2, &cc2); err != nil {
		panic(err)
	}

	fmt.Println(cc2)

}
