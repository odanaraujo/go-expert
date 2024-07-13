package main

import "fmt"

func main() {
	var x interface{} = "Dan Araújo"
	//println(x)          //output: my name is %v (0x2a7040,0x2ccaa0)
	//println(x.(string)) //output: Dan Araújo

	var result string
	var ok bool
	if result, ok = x.(string); !ok {
		println("error")
	}

	fmt.Println(result)
}
