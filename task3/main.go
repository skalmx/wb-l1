package main

import (
	"fmt"
	"sync"
)

func main() {
	array := []int{2, 4, 6, 8, 10}
	w := sync.WaitGroup{}
	m := sync.Mutex{}
	sum := 0

	for _, value := range array { // для того чтобы горутины не имели доступ к счетчику суммы в один момент времени, счетчик обернут в мютекс
		w.Add(1)
		go func(v int) {
			defer w.Done()
			m.Lock()
			sum += v * v
			m.Unlock()
		}(value)
	}
	w.Wait()
	fmt.Println(sum)
}
