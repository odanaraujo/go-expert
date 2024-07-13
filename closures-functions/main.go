package main

import "fmt"

func main() {
	total := func() int {
		return sum(1, 4, 6, 78, 5, 6)
	}()

	fmt.Println(total)
}

func sum(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}

	return total
}
