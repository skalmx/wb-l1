package main

import (
	"fmt"
	"reflect"
)

func main() {
	PrintType(1)
	PrintType("str")
	PrintType(false)
	PrintType(make(chan float64))
	PrintType(1.5)
	
}

func PrintType(value interface{}) {
	fmt.Println(reflect.TypeOf(value))
}