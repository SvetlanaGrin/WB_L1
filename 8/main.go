package main

import (
	"fmt"
)

func main() {
	var i int
	fmt.Print("индекс:")
	fmt.Scanln(&i)
	var flag byte = 0
	fmt.Print("значение:")
	fmt.Scan(&flag)
	if flag == 1 {
		var N int64 = 2
		fmt.Printf("%064b\n", N)
		changeOne(i, N)
	}
	if flag == 0 {
		var N int64
		N = 1<<i + 1<<1 + 1<<2
		fmt.Printf("%064b\n", N)
		changeZero(i, N)
	}

}

func changeOne(i int, N int64) {
	var mask int64
	mask = 1 << i
	c := N ^ mask
	fmt.Printf("%064b\n", c)
}

func changeZero(i int, N int64) {
	var mask int64
	mask = 1 << i
	var mask2 int64
	mask2 = ^mask
	h := N & mask2
	fmt.Printf("%064b\n", h)
}
