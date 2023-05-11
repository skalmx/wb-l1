package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// В данной ситуации можно было работать и с мьютексами, но атомики работают на более низком уровне и дает + в производительности
type Counter struct {
	counter int64
}

func (c *Counter) Increment() {
	atomic.AddInt64(&c.counter, 1) 
}

func main() {
	c := Counter{}

	w := sync.WaitGroup{}

	for i := 0; i < 250330; i++ {
		w.Add(1)
		go func() {
			defer w.Done()
			c.Increment()
		}()
	}
	w.Wait()
	//fmt.Println(c.counter)
	fmt.Println(atomic.LoadInt64(&c.counter))
}