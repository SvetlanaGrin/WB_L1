package main

import "fmt"

func main() {
	arr := []int{1, 5, 5, 5, 12, 14, 20}
	fmt.Println(lbs(arr, 5))
	fmt.Println(rbs(arr, 5))
}
func lbs(arr []int, x int) int {
	l := -1
	r := len(arr) - 1
	for l+1 != r {
		c := (l + r) / 2
		if arr[c] < x {
			l = c

		} else {
			r = c
		}
	}
	return r
}
func rbs(arr []int, x int) int {
	l := 0
	r := len(arr)
	for l != r-1 {
		c := (l + r) / 2
		if arr[c] > x {
			r = c
		} else {
			l = c
		}
	}
	return l
}
