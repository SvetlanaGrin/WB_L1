package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}
	fmt.Println("Exemple1")
	synsMap(arr)
	fmt.Println("Exemple2")
	MyMap(arr)
}

func synsMap(arr []int) {
	wg := sync.WaitGroup{}
	wg.Add(len(arr))
	mSqrt := sync.Map{}
	for _, elem := range arr {
		if _, ok := mSqrt.Load(elem); ok {
			continue
		}
		go func(elem int) {
			defer wg.Done()
			mSqrt.Store(elem, elem*elem*elem)
		}(elem)
	}
	wg.Wait()
	for _, elem := range arr {
		el, _ := mSqrt.Load(elem)
		fmt.Println(el)
	}
}

type Map[K comparable, V any] struct { // Define K,V type parameters
	mu   sync.RWMutex
	data map[K]V
}

func newMap[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{
		data: make(map[K]V),
	}
}
func (m *Map[K, V]) Store(key K, value V) {
	m.mu.Lock()
	m.data[key] = value
	m.mu.Unlock()
}

func (m *Map[K, V]) Load(key K) (V, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	v, ok := m.data[key]
	return v, ok
}

func MyMap(arr []int) {
	wg := sync.WaitGroup{}
	wg.Add(len(arr))
	mSqrt := newMap[int, int]()
	for _, elem := range arr {
		if _, ok := mSqrt.Load(elem); ok {
			continue
		}
		go func(elem int) {
			defer wg.Done()
			mSqrt.Store(elem, elem*elem*elem)
		}(elem)
	}
	wg.Wait()
	for _, elem := range arr {
		el, _ := mSqrt.Load(elem)
		fmt.Println(el)
	}
}
