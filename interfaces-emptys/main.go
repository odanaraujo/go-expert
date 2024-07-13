package main

import "fmt"

func main() {

	var x interface{} = 10
	var y interface{} = "hello world"

	showType(x)
	showType(y)
}

func showType(x interface{}) {
	fmt.Printf("the type of x is %T and the value is %v\n", x, x)
}
