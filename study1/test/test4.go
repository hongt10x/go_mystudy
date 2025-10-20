package main

import "fmt"

func main() {
	var sum int = 0
	//for定义初始值不能用var，可以用局部变量声明方法:=
	for i := 1; i <= 5; i++ {
		sum += i
	}
	fmt.Println(sum)
}
