package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var input int
	fmt.Print("Введите время в секундах, через которое программа закончится:")
	fmt.Scan(&input)

	rand.Seed(time.Now().UnixNano()) // для правильного использования рандомной генерации строк

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * time.Duration(input)) // при завершении времени сработает  cancel у контекста и завершиться работа программы
	defer cancel()

	ch := make(chan string)

	go func(context.Context, chan string) { 
		for {
			select {
			case <- ctx.Done(): // если время выйдет, данные перестанут писаться 
				close(ch)
				return
			default:
				ch <- RandString(10) //пишем в канал рандомную строку
				time.Sleep(time.Second)
			}
		}
	}(ctx, ch)

	for value := range ch {
		fmt.Println("Info:", value)
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