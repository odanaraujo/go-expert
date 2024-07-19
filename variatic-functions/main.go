package main

import "fmt"

func main() {

	fmt.Println(sum(20, 10, 45, 22, 4, 5))
}

func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}

	return total
}
