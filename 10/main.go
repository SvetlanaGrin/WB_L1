package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	temp := []int{-11, -2, 14, 23, 37, 24, -25, -32, 5, 8}
	mapTemp := make(map[string]string)
	strBuilder := strings.Builder{}
	for _, elem := range temp {
		strBuilder.WriteString(mapTemp[strconv.Itoa((elem/10)*10)] + ", " + strconv.Itoa(elem))
		mapTemp[strconv.Itoa((elem/10)*10)] = strBuilder.String()
		strBuilder.Reset()
	}

	for i := -50; i <= 50; i = i + 10 {
		if mapTemp[strconv.Itoa(i)] != "" {
			strBuilder.WriteString(strconv.Itoa(i) + ":{" + mapTemp[strconv.Itoa(i)][2:] + "}, ")
		}

	}
	strBuilder.WriteString("etc.")
	fmt.Println(strBuilder.String())
}
