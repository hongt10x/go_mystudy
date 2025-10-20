package main

//匿名函数
import (
	"fmt"
)

func main() {
	//匿名函数
	ret := func(num1, num2 int) int {
		fmt.Println(num1, num2)
		return num1 + num2
	}(10, 20)

	fmt.Println(ret)
	num1 := 100
	fmt.Println(num1)
	fmt.Println(ret)

	ret1 := func(num1, num2 int) int {
		return num1 + num2
	}

	fmt.Println(ret1(30, 40))
}
