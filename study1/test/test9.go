package main

import (
	"fmt"
)

func main() {
	fmt.Println(f1(10, 20))
}

func f1(num11, num21 int) int {
	//defer将该行语句压入栈中，先进后出
	defer fmt.Println("num11", num11)
	defer fmt.Println("num21", num21)
	fmt.Println("******")
	num11 += 90
	num21 += 50
	var num31 int = num11 + num21
	fmt.Println("num31", num31)
	return num31
}
