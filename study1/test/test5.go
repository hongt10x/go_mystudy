package main

import "fmt"

func main() {
	//中文占3个字节
	var str string = "hello golang,你好"
	//for定义初始值不能用var，可以用局部变量声明方法:=
	//for i := 0; i < len(str); i++ {
	//	fmt.Printf("%c\n", str[i])
	//}

	//方式二
	//for i, value := range str {
	//	fmt.Println(i, string(value))
	//}

	//for i, _ := range str {
	//	fmt.Println(i)
	//}

	//for _, value := range str {
	//	fmt.Println(value, string(value))
	//}

	//for k := range str {
	//	fmt.Println(k)
	//}
	for k, v := range str {
		fmt.Println(k, string(v))
	}
}
