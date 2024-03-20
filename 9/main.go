package main

import (
	"fmt"
	"time"
)

func producer(nums []int, numbers chan int) {
	for _, elem := range nums {
		numbers <- elem
		time.Sleep(time.Second)
	}
}

func square(numbers <-chan int, squares chan<- int) {
	for num := range numbers {
		squared := num * 2
		squares <- squared
	}
}

func consumer(squares <-chan int) {
	for squared := range squares {
		fmt.Println("Received:", squared)
	}
}

func main() {
	numbers := make(chan int)
	squares := make(chan int)
	go producer([]int{3, 6, 14, 3, 566, 4, 13, 45, 8}, numbers)
	go square(numbers, squares)
	go consumer(squares)

	// Чтобы программа не завершилась сразу, добавим ожидание
	time.Sleep(10 * time.Second)
}
