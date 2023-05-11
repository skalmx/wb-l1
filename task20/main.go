package main

import (
	"fmt"
	"strings"
)
// В данном задании, строка разделяется по пробелам и в цикле элементы входящей строки меняются местами
func main() {
	input := "test case - reverse string"

	fmt.Println(Reverse(input))
}

func Reverse(input string) string {
	temp := strings.Split(input, " ")

	for i, j := 0, len(temp)-1; i < j; i, j = i+1, j-1 {
		temp[i], temp[j] = temp[j], temp[i]
	}

	result := strings.Join(temp, " ")

	return result
}