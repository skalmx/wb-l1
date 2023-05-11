package main

import "fmt"

func main() {
	array := []string{"cat", "cat", "dog", "cat", "tree"}
	PrintNewSet(array)
}

func PrintNewSet(array []string) {
	set := make(map[string]struct{})

	for _, value := range array {
		set[value] = struct{}{}
	}

	for value := range set {
		fmt.Println(value)
	}
}