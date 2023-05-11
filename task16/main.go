package main

import "fmt"

func main() {
	array := []int{0, -120, 1, 2424, 2, -5, 3}
	fmt.Println(QuickSort(array, 0, len(array) - 1))
}

func Partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func QuickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = Partition(arr, low, high)
		arr = QuickSort(arr, low, p-1)
		arr = QuickSort(arr, p+1, high)
	}
	return arr
}