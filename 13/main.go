package main

import "fmt"

func main() {
	a := 15
	b := 34

	a, b = b, a
	fmt.Println(a, b)

	a, b = 15, 34

	a = a * b
	b = a / b
	a = a / b
	fmt.Println(a, b)

	a, b = 15, 34
	a = a + b
	b = a - b
	a = a - b
	fmt.Println(a, b)

}
