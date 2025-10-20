package main

import "fmt"

func t1(num ...int) {
	sum := 0
	for _, v := range num {
		sum += v
	}
	fmt.Println(sum)
}

func Print2[T any](slice []T) {
	for _, v := range slice {
		fmt.Println(v)
	}
}

func main() {
	// 多个入参输入，类似与python中的args
	t1(1, 3, 5, 6, 7)
	// 入参为切片
	number := []int{1, 2, 3}
	t1(number...) // 使用 ... 将切片展开为可变参数

	intSlice := []int{1, 2, 3}
	stringSlice := []string{"a", "b", "c"}

	Print2(intSlice)
	Print2(stringSlice)

}
