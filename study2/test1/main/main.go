package main

import (
	"fmt"

	d1 "study2/test1/dbutil"
)

func main() {

	fmt.Println("我是main.go中的方法")
	fmt.Println("我调用了good.go中的属性SCHOOL:", d1.SCHOOL)
	fmt.Println("我开始调用good.go中的GetCount方法，返回:", d1.GetCount())

	//SayHi()
	//fmt.Println(name)

}
