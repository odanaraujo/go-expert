package main

import "fmt"

type person interface {
	inactive()
}

type people struct {
	name   string
	age    int
	status bool
}

func (p people) inactive() {
	p.status = false
	fmt.Printf("people %s is inactive", p.name)
}

func inactived(p people) {
	p.inactive()
}

func main() {
	dan := people{
		name:   "Dan",
		age:    34,
		status: true,
	}

	fmt.Println(dan)

	dan.inactive()

	//eu posso passar o dan no parâmetro, mesmo sendo do tipo pessoa interface, pois people implementa essa interface
	inactived(dan)
}
