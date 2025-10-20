package main

import (
	"fmt"
	"time"
)

func printNum() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

// 捕获错误用defer+recover
func sub() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	num1, num2 := 10, 0
	fmt.Println(num1 / num2)
}

func main() {
	go printNum()
	go sub()
	//fmt.Println(time.Second)
	time.Sleep(time.Second * 3)
}
