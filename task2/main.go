package main

import (
	"fmt"
	"sync"
)

func main() {
	array := []int{2, 4, 6, 8, 10}
	w := sync.WaitGroup{} 

	for _, value := range array {
		w.Add(1)
		go func(v int) {
			defer w.Done()
			fmt.Printf("%d в квадрате = %d\n", v, v * v)
		}(value)
	}
	w.Wait()
}