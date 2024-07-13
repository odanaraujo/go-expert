package main

import "fmt"

type employes struct {
	name   string
	active bool
}

func main() {
	dan := employes{
		name:   "Dan",
		active: true,
	}

	sabrina := employes{
		name:   "Sabrina",
		active: false,
	}
	salarys := map[employes]float32{dan: 100, sabrina: 200}

	for indice, value := range salarys {
		if !indice.active {
			delete(salarys, indice)
			break
		}
		fmt.Printf("salary of %v is %.2f\n", indice, value)
	}

	//uma segunda forma de criar map é usando a função make
	sal := make(map[employes]float32)
	fmt.Printf("the type %T", sal)
}
