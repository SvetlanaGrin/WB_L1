package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}
	wg.Add(len(arr))
	mSqrt := sync.Map{}
	for _, elem := range arr {
		if _, ok := mSqrt.Load(elem); ok {
			continue
		}
		go func(elem int) {
			defer wg.Done()
			mSqrt.Store(elem, elem*elem)
		}(elem)
	}
	wg.Wait()
	for _, elem := range arr {
		el, _ := mSqrt.Load(elem)
		fmt.Println(el)
	}

}
