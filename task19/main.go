package main

import "fmt"

func main() {
	fmt.Println(Reverse("главрыба"))
}

func Reverse(word string) (result string) {
	for _, value := range word {
		result = string(value) + result // здесь происходит переприсваивание изначальной строки в строку, где добавлена новая буква + старая строка
	}
	return result
}
