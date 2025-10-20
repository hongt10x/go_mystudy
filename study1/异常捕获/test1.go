package main

import (
	"errors"
	"fmt"
)

func main() {
	test()
	err := test2()
	fmt.Println(err)
	test()

	fmt.Println("test函数执行完毕")
	fmt.Println("逻辑正常开始执行下面的函数")
}

func test() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	num1 := 10
	num2 := 0
	result := num1 / num2
	fmt.Println(result)
}

func test2() error {

	num1 := 10
	num2 := 0
	if num2 == 0 {
		return errors.New("除数不能为0 ~_~")
	} else {
		result := num1 / num2
		fmt.Println(result)
		return nil
	}
}
