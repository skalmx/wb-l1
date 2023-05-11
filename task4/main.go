package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var input int
	fmt.Println("Введите число")
	fmt.Scan(&input) // получаем количество воркеров

	rand.Seed(time.Now().UnixNano()) // для правильного использования рандомной генерации строк

	ctx, cancel := context.WithCancel(context.Background())
	quit := make(chan os.Signal, 1) // перехватывает систменые сигналы и передаем их в канал 
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	ch := make(chan string)

	StartWorkers(ctx, input, ch)
	
	for {
		select {
		case <- quit: // при получении системных сигналов завершаем работу
			cancel()
			return
		default: 
			ch <- RandString(50) // передаем рандомную строку из 5 символов в канал
		}	
	}
}

func StartWorkers(ctx context.Context, n int, ch chan string) {
	for i := 0; i < n; i++ {
		go func(i int) {
			for {
				select {
				case <- ctx.Done():
					return
				case value := <- ch:
					fmt.Printf("Worker №%d Said:%s\n", i, value)
				}
			}
		}(i)	

	}
}

func RandString(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
    b := make([]rune, n)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return string(b)
}