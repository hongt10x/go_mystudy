package main

import "fmt"

type SayHello interface {
	sayHello()
}

type Chinese struct {
	name string
}

func (p Chinese) sayHello() {
	fmt.Println("你好，我是中国人")
}

func (p Chinese) niuYangGe() {
	fmt.Println("中国人会扭秧歌")
}

type American struct {
	name string
}

func (p American) sayHello() {
	fmt.Println("hello， I am a American")
}

func greet(s SayHello) {
	s.sayHello()
	//s.niuYangGe()
}

func main() {
	c := Chinese{}
	greet(c)
	a := American{}
	greet(a)

}
