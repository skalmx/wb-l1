package main

import (
	"fmt"
	"sync"
)

func main() {
	testMap := NewCustomMap()
	testMap.Write("1", 1)
	testMap.Write("2", 22222)
	fmt.Println(testMap.Get("1"), testMap.Get("2"))
}

type CustomMap struct { // с помощью мютекса в среде с конкрутным доступом горутины будут получать блокировку на запись, но смогут одновременно читать данные из мапы
	sync.RWMutex
	data map[string]interface{}
}

func NewCustomMap() *CustomMap {
	return &CustomMap{
		data: make(map[string]interface{}),
	}
}

func (c *CustomMap) Write(key string, value interface{}) {
	c.Lock()
	defer c.Unlock()
	c.data[key] = value
}

func (c *CustomMap) Get(key string) interface{} {
	c.Lock()
	defer c.Unlock()
	return c.data[key]
}