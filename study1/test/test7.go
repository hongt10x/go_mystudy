package main

import "fmt"

func calc(x, y int) int {
	return x + y
}

func test1(x, y int) (int, int) {
	sum1 := x + y
	sub1 := x - y
	return sum1, sub1
}

//func test2(x int, y int) (sub1 int, sum1 int) {
//	sum1 := x + y
//	sub1 := x - y
//	return
//}

func init() {
	fmt.Println("init函数被调用")
}

func test3() int {
	fmt.Println("test函数被调用")
	return 99
}

var num int = test3()

func main() {
	fmt.Println("main函数被调用")
	//res := calc(3, 5)
	//fmt.Println(res)
	fmt.Println(num)
	//z1, z2 := test1(3, 5)
	//fmt.Println(z1, z2)
	//z3, z4 := test2(3, 5)
	//fmt.Println(z3, z4)
	fmt.Println("-------------------")
	var num1 int = test3()
	fmt.Println(num1)
}
