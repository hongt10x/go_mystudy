package main

import "fmt"

/*
闭包引入的外部参数可以持久化使用，不洗白，不清零
*/

func GetSum() func(int) int {
	var sum int = 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	f := GetSum()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))
}
