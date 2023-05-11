package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	Sleep(3)
	fmt.Println(time.Now())
}

func Sleep(x int) {
	<-time.After(time.Second * time.Duration(x))
}
