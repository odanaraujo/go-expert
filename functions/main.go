package main

import "fmt"

func main() {
	fmt.Println(sum(2, 3))
}

//podemos retornar mais de um valor na function
// se o valor passado via parâmetro for do mesmo tipo, podemos colocar o nome das variáveis e o tipo no final.
func sum(a, b int) (int, bool) {
	if a+b >= 50 {
		return a + b, true
	}
	return a + b, false
}
