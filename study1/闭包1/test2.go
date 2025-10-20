package main

import "fmt"

func GetSum1() func(int) int {
	var sum int = 0
	return func(i int) int {
		sum += i
		return sum
	}
}

func main() {
	f := GetSum1()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))
}
