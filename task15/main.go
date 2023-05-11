package main

import (
	"fmt"
	"strings"
)

var justString string

func someFunc() {
	v := createHugeString(1 << 10) // в данном случае если символы будут занимать вместо 1 байта - 2 (к примеру русские буквы), будем получать половину строки (50 символов вместо 100, так как 1 символ теперь будет занимать два байта)
	justString = v[:100]

}

func main() {
	someFunc()
	fmt.Println(justString) // 50 символов 
	someFuncFix()
	fmt.Println(justString) // 100 cимволов как и нужно
}

func createHugeString(n int) string {
	str := strings.Builder{}

	for i := 0; i < n; i++ {
		str.WriteString("а")
	}

	return str.String()

}

func someFuncFix() {
	v := []rune(createHugeString(1 << 10))
	justString = string(v[:100])
}
