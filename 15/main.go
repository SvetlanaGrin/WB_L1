package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	v := createHugeString(1 << 10)

	justString := string(CopyBytes([]rune(v), 100))
	fmt.Println(justString)
}

func createHugeString(n int) string {
	letter := "dfghbfldkjbdhvhаваплвиомариялвоми"
	runes := []rune(letter)
	var str strings.Builder
	str.Grow(n)
	for i := 0; i < n; i++ {
		str.WriteRune(runes[rand.Intn(len(runes))])
	}
	return str.String()
}

func CopyBytes(input []rune, n int) []rune {
	result := make([]rune, n)
	for i := 0; i < n; i++ {
		result[i] = input[i]
	}
	return result
}
