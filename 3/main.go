package main

import (
	"fmt"
	"sync"
)

type Sum struct {
	mu   sync.RWMutex
	data int
}

func NewSum(data int) *Sum {
	return &Sum{data: data}
}
func (s *Sum) addSum(input int) {
	s.mu.Lock()
	s.data = s.data + input
	s.mu.Unlock()
}
func (s *Sum) getSum() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.data
}
func main() {

	arr := []int{2, 4, 6, 8, 10, 12}
	N := len(arr)
	ch := make(chan int, N)
	defer close(ch)
	for _, elem := range arr {
		ch <- elem
	}
	sum := NewSum(0)
	wg := sync.WaitGroup{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func() {
			defer wg.Done()
			elem := <-ch
			sum.addSum(elem * elem)
		}()
	}
	wg.Wait()
	fmt.Println(sum.getSum())

}
