package main

import "fmt"

// 复杂
func tet(func1 func(int, int) int, num1 int, num2 int) int {
	return func1(num1, num2)
}

// 函数别名 funcTestType
type funcTestType func(int, int) int

// 简单方法
func tet2(func1 funcTestType, num1 int, num2 int) int {
	return func1(num1, num2)
}

func getSum(a, b int) int {
	return a + b
}

func main() {
	a := tet(getSum, 1, 2)
	fmt.Println(a)

	b := tet2(getSum, 1, 2)
	fmt.Println(b)

	type myInt int
	var a1 myInt = 100
	fmt.Println(a1)
}

/*
package calc


import "fmt"

func test1() {
	fmt.Println("hello world")
}

type t1 test1

func calc() {
	t1()
}
*/
