package main

import (
	"fmt"
	"math/big"
)

func main() {
	x := new(big.Int)
	x.SetString("12000000000000000000", 10)
	y := new(big.Int)
	y.SetString("2400000000000000000000", 10)
	z := new(big.Int)
	fmt.Println("Add: ", z.Add(x, y))
	fmt.Println("Sub: ", z.Sub(x, y))
	fmt.Println("Mul: ", z.Mul(x, y))
	fmt.Println("Div: ", z.Div(y, x))

}
