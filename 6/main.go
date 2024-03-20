package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// создаем контекст с отменой
	ctx, cancel := context.WithCancel(context.Background())
	// запускаем горутину с этим контекстом

	go workerStopCtx(ctx)
	stop := make(chan bool)
	// запускаем горутину с этим каналом
	go workerStopChan(stop)

	// ждем 3 секунды
	time.Sleep(2 * time.Second)
	// отменяем контекст
	cancel()
	stop <- true
	// ждем еще 2 секунды
	time.Sleep(1 * time.Second)
}

func workerStopCtx(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			// контекст отменен, выходим из горутины
			fmt.Println("worker 1 stopped")
			return
		default:
			// продолжаем работать
			fmt.Println("worker 1")
			time.Sleep(time.Second)
		}
	}
}
func workerStopChan(stop chan bool) {
	for {
		select {
		case <-stop:
			// получили сигнал остановки, выходим из горутины
			fmt.Println("worker 2 stopped")
			return
		default:
			// продолжаем работать
			fmt.Println("worker 2")
			time.Sleep(time.Second)
		}
	}
}
