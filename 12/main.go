package main

import (
	"fmt"
)

// структура множество
type Set[T comparable] struct {
	data map[T]struct{}
}

func NewSet[T comparable]() Set[T] {
	return Set[T]{
		data: make(map[T]struct{}),
	}
}

// метод добавления значений в множество
func (s *Set[T]) Add(values ...T) {
	for _, val := range values {
		if _, ok := s.data[val]; !ok {
			s.data[val] = struct{}{}
		}
	}
}
func (s *Set[T]) Print() {
	fmt.Printf("Set{ ")
	for val := range s.data {
		fmt.Print(val, " ")
	}
	fmt.Print("}")
}

func main() {
	set := NewSet[string]()
	set.Add("cat", "cat", "dog", "cat", "tree")
	set.Print()

}
