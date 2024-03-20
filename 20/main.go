package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	str := "snow dog sun"
	fmt.Println(reverse(str))
	fmt.Println(reverse2(str))
}

func reverse(input string) string {

	arrInput := strings.Split(input, " ")
	var output string
	for i := 0; i < len(arrInput); i++ {
		output = " " + arrInput[i] + output
	}
	return strings.TrimSpace(output)
}

func reverse2(input string) string {

	arrInput := strings.Split(input, " ")
	slices.Reverse(arrInput)
	return strings.Join(arrInput, " ")
}
