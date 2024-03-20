package main

import (
	"fmt"
)

func getType(input interface{}) {
	switch input.(type) {
	case string:
		fmt.Printf("%s ,type(%T)\n", input, input)
	case int:
		fmt.Printf("%d ,type(%T)\n", input, input)
	case bool:
		fmt.Printf("%t ,type(%T)\n", input, input)
	case chan int:
		fmt.Printf("%p ,type(%T)\n", input, input)

	}
}

func main() {
	getType("main")
	getType(11)
	getType(true)
	getType(make(chan int))
	var input interface{}
	input = 123
	fmt.Printf("%d ,type(%T)\n", input.(int), input.(int))
}
