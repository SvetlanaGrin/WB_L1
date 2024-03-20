package main

import (
	"fmt"
	"strings"
)

func main() {

	str := string("abcderTUt")
	fmt.Println("func1", unique(str))
	fmt.Println("func2", unique2(str))
	str = string("abcde")
	fmt.Println("func1", unique(str))
	fmt.Println("func2", unique2(str))

}

func unique(input string) bool {
	input = strings.ToLower(input)
	mapInput := make(map[int32]int)
	for _, elem := range input {
		mapInput[elem] += 1
	}
	for key, _ := range mapInput {
		if mapInput[key] > 1 {
			return false
		}
	}
	return true
}
func unique2(input string) bool {
	input = strings.ToLower(input)
	for _, elem := range input {
		if strings.Count(input, string(elem)) != 1 {
			return false
		}
	}
	return true
}
