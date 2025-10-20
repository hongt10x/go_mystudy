package main

import "fmt"

var NAME = "wht"
var name = "echo"

func SayHi() {
	fmt.Println("我是good.go中的方法，我有个属性NAME=", NAME)
	fmt.Println("我还有个隐私小名字name=", name)
}

//func calc() {
//	SayHi()
//	fmt.Println(SayHi())
//}
