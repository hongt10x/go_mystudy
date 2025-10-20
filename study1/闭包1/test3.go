package main

import (
	"errors"
	"fmt"
)

type operate func(x, y int) int

type calculateFunc func(x int, y int) (int, error)

func genCalculator(op operate) calculateFunc {
	return func(x int, y int) (int, error) {
		if op == nil {
			return 0, errors.New("invalid operation")
		}
		return op(x, y), nil
	}
}

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
