package main

import (
	"fmt"
	"math/big"
)

func main() {
	x := big.NewInt(3000000)
	y := big.NewInt(2000001)

	sum := new(big.Int).Add(x, y)
	fmt.Println("Сложение = ", sum)
	sub := new(big.Int).Sub(x, y)
	fmt.Println("Вычитание = ", sub)
	mul := new(big.Int).Mul(x, y)
	fmt.Println("Умножение = ", mul)
	div := new(big.Int).Quo(x, y)
	fmt.Println("Деление = ", div)

}