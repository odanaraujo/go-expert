package main

import "fmt"

type client struct {
	name string
}

type account struct {
	value int
}

func (a *account) simulation(value int) int {
	a.value += value
	fmt.Printf("the balance is %d\n", a.value)
	return a.value
}

func (c client) walked() {
	c.name = "Dan Araújo"
	fmt.Printf("client %s walked\n", c.name)
}

func main() {

	dan := client{
		name: "Dan",
	}

	dan.walked()
	fmt.Printf("value name in struct is %v\n", dan)

	//exemplo com conta
	accountDan := account{
		value: 240,
	}

	accountDan.simulation(120)
	fmt.Printf("the value in struct is %v", accountDan)
}
