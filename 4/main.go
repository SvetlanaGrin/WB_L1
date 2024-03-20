package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Print(id, ":")
		fmt.Println(j)
	}
}

func main() {
	const numJobs = 5
	const workersCount = 3
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	wg := sync.WaitGroup{}
	wg.Add(workersCount)
	jobs := make(chan int, numJobs)

	for w := 1; w <= workersCount; w++ {
		go worker(w, jobs, &wg)
	}
	for {
		select {
		case <-c:
			close(jobs)
			wg.Wait()
			return
		default:
			jobs <- rand.Intn(100)
		}
	}

}
