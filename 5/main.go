package main

import (
	"fmt"
	"time"
)

func main() {
	var N int64
	fmt.Scanln(&N)
	timeout := time.After(time.Second * time.Duration(N))
	i := 0
	ch := make(chan int)
	go func() {
		for {
			select {
			case <-time.After(time.Second * time.Duration(1)):
				ch <- i
				i++
			case <-timeout:
				fmt.Println(" out: ", i)
				close(ch)
				return
			}
		}
	}()
	for j := range ch {
		fmt.Println("value: ", j)
	}
}
