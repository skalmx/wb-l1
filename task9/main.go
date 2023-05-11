package main

import "fmt"

func main() {
	array := []int{2, 4, 6, 8, 10}
	//ch := Write(array)
	result := Square(Write(array))
	for value := range result {
		fmt.Println(value)
	}
}

func Square(ch chan int) chan int {
	sq := make(chan int) // создаем канал для квадратов x
	go func() {
		defer close(sq)
		for value := range ch {
			sq <- value * value
		}
	}()

	return sq
}

func Write(input []int) chan int {
	ch := make(chan int) // создаем канал для x
	go func() {
		defer close(ch)
		for _, value := range input {
			ch <- value
		}
	}()

	return ch
}
