package main

import (
	"fmt"
	"strings"
)

func main() {
	test := "aAbcdefg"
	fmt.Println(CheckForUnique(test)) //false
	test = "abcdef"
	fmt.Println(CheckForUnique(test)) //true
}

func CheckForUnique(str string) bool {
	set := make(map[rune]struct{}) // храним уникальные символы в мапе

	lowString := strings.ToLower(str) // приводим строку к меньшому регистру

	for _, value := range lowString {
		if _, exists := set[value]; exists {
			return false
		} else {
			set[value] = struct{}{}
		}
	}
	return true
}