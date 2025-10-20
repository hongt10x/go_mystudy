package main

import "fmt"

func main() {
	fmt.Println("main函数被执行")
	fmt.Println("name的值：", name)
}

func init() {
	fmt.Println("init函数被执行")
}

var name string = t1()

func t1() string {
	fmt.Println("t1函数被调用了")
	return "wang"
}

/*  输出结果不因为位置的变化而变化
t1函数被调用了
init函数被执行
main函数被执行
name的值： wang
*/
