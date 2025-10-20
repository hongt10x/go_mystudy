package main

import (
	"fmt"
)

func main() {

	fmt.Println("HELLO, Go")
	var age int
	fmt.Print("请用户输入年龄-->")
	fmt.Scanln(&age)
	fmt.Printf("当前用户的年龄是%d岁\n", age)
	var (
		//age   int
		name  string
		score float32
		isVIP bool
	)
	fmt.Println("请录入学生年龄，多个以空格分割")
	fmt.Scanf("%s %d %f %t", &name, &age, &score, &isVIP)
	fmt.Printf("学生姓名：%v,年龄%v岁,考了%v分,是否是VIP:%v \n", name, age, score, isVIP)
	fmt.Printf("学生姓名：%#v,年龄%#v岁,考了%#v分,是否是VIP:%#v\n", name, age, score, isVIP)
	fmt.Printf("学生姓名：%+v,年龄%+v岁,考了%+v分,是否是VIP:%+v\n", name, age, score, isVIP)
}
