package main

import "fmt"

func main() {
	var myArray [3]int //array de tamanho fixo. Não pode colocar mais que três valores
	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 3

	// último elemento = 2
	fmt.Println(len(myArray) - 1)

	//imprime o valor do último item
	fmt.Println(myArray[len(myArray)-1])

	//for com range onde o primeiro parâmetro é o index e o segundo o valor
	for index, value := range myArray {
		fmt.Printf("position %d has the value %d\n", index, value)
	}
}
