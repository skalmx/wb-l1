package main

import (
	"fmt"
	"sort"
)

func main(){
	array := []int{1, 2, 3, 4, 5, 6, 1233, 23232, 23, 68, 63, 102}
	sort.Ints(array)
	fmt.Println(BinarySearch(1233, array))
}

func BinarySearch(value int, list []int) string {

	low := 0
	high := len(list) - 1

	for low <= high{
		mid := (low + high) / 2
		if list[mid] > value {
			high = mid - 1
		} else if list[mid] < value {
			low = mid + 1
		} else {
			return "Found your number"
		}
	}

	return  "Dont found your number"

}
