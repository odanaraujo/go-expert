package main

import "fmt"

func main() {
	s := []int{2, 4, 6, 8, 10}

	fmt.Printf("len=%d, cap=%d, %v\n", len(s), cap(s), s)

	//a partir da posição 0, tudo que estiver a direita é para dropar
	fmt.Printf("len=%d, cap=%d, %v\n", len(s[:0]), cap(s[:0]), s[:0])

	fmt.Printf("len=%d, cap=%d, %v\n", len(s[:4]), cap(s[:4]), s[:4])

	//A partir da posição 2, tudo que estiver à esquerda vai ser dropado
	//Uma observação aqui é em relação ao cap, que será diminuído de 5 para 3.
	fmt.Printf("len=%d, cap=%d, %v\n", len(s[2:]), cap(s[2:]), s[2:])

	s = append(s, 12)

	//por debaixo dos panos, o slice dobrou a capacidade de 5 para 10.
	fmt.Printf("len=%d, cap=%d, %v\n", len(s), cap(s), s)
	fmt.Printf("len=%d, cap=%d, %v\n", len(s[2:]), cap(s[2:]), s[2:])
}
