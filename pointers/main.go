package main

import "fmt"

func main() {
	a := 10
	var pointer *int = &a
	*pointer = 20
	fmt.Println(*pointer)
	b := &a
	*b = 30
	fmt.Println(*b)

	// quando usar ponteiro?
	// Quando não queremos passar uma cópia do valor.

	number1 := 10
	number2 := 20
	fmt.Println(sum(number1, number2)) // a soma vai ser 70, pois dentro do sum eu altero o a para 50.
	fmt.Println(number1)               // vai imprimir 10 pois eu alterei a cópia do a e não a memória

	var numberPointer1 *int = &number1
	var numberPointer2 *int = &number2

	fmt.Println(sum2(numberPointer1, numberPointer2)) // a soma será de 60, pois altero o a dentro do sum2
	fmt.Println(*numberPointer1)                      // o output será 40, pois nesse momento alterei o valor na memória e não a cópia

	//outra forma de passar o endereço da memória
	fmt.Println(sum2(&number1, &number2))
	fmt.Println(*&number1)
}

func sum(a, b int) int {
	a = 50 // estou alterando o valor da cópia de number1, mas number1 continua com o valor que foi criado
	return a + b
}

func sum2(a, b *int) int {
	*a = 40
	return *a + *b
}
