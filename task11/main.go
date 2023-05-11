package main

import "fmt"

func main() {
	array1 := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	array2 := []interface{}{1, 3, 5, 7, 9, 11, 13}
	PrintIntersection(array1, array2) // 1 3 5 7 9 11

	array3 := []interface{}{"test", "not-test","something"}
	array4 := []interface{}{"test1", "not-test", "test"}
	PrintIntersection(array3, array4) // not-test test
}

func PrintIntersection(array1 []interface{}, array2 []interface{}) { // чтобы находить пересечения множеств разных типов решил использовать пустые интерфейсы
	check := make(map[interface{}]struct{})
	var result []interface{}

	for _, value := range array1 {
		check[value] = struct{}{}
	}

	for _, value := range array2 {
		if _, exists := check[value]; exists {
			result = append(result, value)
		}
	}

	fmt.Println(result...)
}