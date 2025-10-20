package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	var str1 string = "hello world 中国！"
	fmt.Println(len(str1))
	//
	//for i, value := range str1 {
	//	fmt.Println(i, string(value))
	//}

	//r := []rune(str1)
	//for i := 0; i < len(r); i++ {
	//	fmt.Printf("%c\n", r[i])
	//}
	//将字符串转为数字 atoi (表示 ascii to integer)是把字符串转换成整型数的一个函数，应用在计算机程序和办公软件中
	num1, err := strconv.Atoi("666")
	fmt.Println(num1, err)
	//统计指定字符串的数量
	count := strings.Count(str1, "hello")
	fmt.Println(count)
	//不区分大小写
	flag := strings.EqualFold("HELLO", "hello")
	fmt.Println(flag)
	//区分大小写
	fmt.Println("HELLO" == "hello")
	//从第几个开始
	fmt.Println(strings.Index(str1, "llo"))
}
