package main

import (
	"fmt"
	"slices"
)

func main() {
	str := "главрыба"
	fmt.Println(reverse(str))
	fmt.Println(reverse2(str))
	fmt.Println(reverse3(str))
}

func reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}

func reverse2(s string) string {
	var str string
	for _, v := range s {
		str = string(v) + str
	}
	return str
}
func reverse3(s string) string {
	rns := []rune(s)
	slices.Reverse(rns)
	return string(rns)
}
