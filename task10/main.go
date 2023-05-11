package main

import "fmt"

func main() {
	array := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	result := make(map[int][]float32)

	for _, value := range array {
		key := int(value / 10) * 10
		result[key] = append(result[key], value)
	}
	fmt.Print(result)
}