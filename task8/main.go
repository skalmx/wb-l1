package main

import "fmt"

func main() {
	var test int64
	test = 11
	fmt.Printf("test - %b\n", test) // 1011
	fmt.Printf("test after changing 0 to 1 - %b\n",ChangeTOne(test, 3)) // 1111
	fmt.Printf("test after changing 1 to 0 - %b\n",ChangeToZero(test, 1))//1010
}

func ChangeTOne(number int64, pos int) int64 { // тут происходит операция побитового или нашего исходного числа и 1 сдвинутой влево на заданную позицию
	return number | 1 << (pos - 1)
}

func ChangeToZero(number int64, pos int) int64 { // тут уже происходит операция логического отрицания 1 сдвинутой на заданную позицию и операция побитового и
	return number &^(1 << (pos - 1))
}

