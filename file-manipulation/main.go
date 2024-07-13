package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("file.txt")
	if err != nil {
		panic(err)
	}

	//lenght, err := file.WriteString("Hello, World") //método salva string, mas podemos querer salvar bytes
	lenght, err := file.Write([]byte("Sport club do Recife é o maior do norte/nordeste"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("File created successfully! Tamanho: %d bytes\n", lenght)
	readFile, err := os.ReadFile("file.txt")

	if err != nil {
		panic(err)
	}

	fmt.Printf("value in file is: %s\n", string(readFile))

	//leitura de pouco em pouco abrindo o arquivo

	readOpen, err := os.Open("file.txt")
	if err != nil {
		panic(err)
	}

	//utilizaremos o bufferio
	reader := bufio.NewReader(readOpen)
	buffer := make([]byte, 10)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		value := string(buffer[:n])
		fmt.Printf("%s", value)
	}
	readOpen.Close()
	file.Close()

	if err := os.Remove("file.txt"); err != nil {
		panic(err)
	}
}
