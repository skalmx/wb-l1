package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {
	// Завершение горутины при помощи контекста
	w := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())
	w.Add(1)
	go func() {
		defer w.Done()
		<-ctx.Done() 
		fmt.Println("Завершение горутины с помощью context")
	}()
	cancel()
	// Завершение горутины с помощью чтения каких либо данных из канала
	quit := make(chan struct{})
	w.Add(1)
	go func() {
		defer w.Done()
		<-quit
		fmt.Println("Завершение горутины с помощью каналов")
	}()
	quit <- struct{}{}

	w.Wait()

	// Также можно завершать горутину через runtime.Goexit()
}