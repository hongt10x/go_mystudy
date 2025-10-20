package main

import (
	"fmt"
)

type operate1 func(x ...int) int

type calculateFunc1 func(x ...int) (int, error)

func main() {
	x, y := 3, 4
	op := func(x, y int) int {
		return x + y
	}
	add := genCalculator(op) // 加法
	result, err := add(x, y)
	fmt.Printf("The addition result: %d (error: %v)\n",
		result, err)

	op1 := func(x, y int) int {
		return x * y
	}
	multi := genCalculator(op1) //乘法
	result, err = multi(x, y)
	fmt.Printf("The multiplication result: %d (error: %v)\n",
		result, err)
}
