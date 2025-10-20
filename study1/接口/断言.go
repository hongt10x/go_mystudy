package main

import "fmt"

type SayHello1 interface {
	sayHello()
}

type Chinese1 struct {
	name string
}

// 通过将sayHello()方法添加到Chinese和American结构体中，这两个结构体都隐式地实现了SayHello接口。在Go中，不需要显式声明“我实现了这个接口”，只要类型拥有接口所需的所有方法，就认为它实现了该接口。

func (p Chinese1) sayHello() {
	fmt.Println("你好，我是中国人")
}

func (p Chinese1) niuYangGe() {
	fmt.Println("中国人会扭秧歌")
}

type American1 struct {
	name string
}

func (p American1) sayHello() {
	fmt.Println("hello， I am a American")
}

func (p American1) disco() {
	fmt.Println("American can dance disco")
}

func greet1(s SayHello1) {
	s.sayHello()
	//断言：写法1
	ch, flag := s.(Chinese1)
	fmt.Printf("ch=%T, flag=%v\n", ch, flag)
	////ch.niuYangGe()
	//if flag {
	//	fmt.Printf("当前对象的类型为%T\n", ch)
	//	ch.niuYangGe()
	//}

	//写法2
	//fmt.Println("------------------------------")
	//if ch1, flag1 := s.(Chinese1); flag1 {
	//	fmt.Printf("当前对象的类型为%T\n", ch1)
	//	//fmt.Println(s.(type))
	//	ch1.niuYangGe()
	//}

	//写法3
	//fmt.Println("------------------------------")
	//switch s.(type) {
	//case Chinese1:
	//	Chinese1{}.niuYangGe()
	//case American1:
	//	//fmt.Println("美国人不会扭秧歌")
	//	American1{}.disco()
	//}

	//写法4
	fmt.Println("------------------------------")
	//switch s.(type) {
	//case Chinese1:
	//	ch1 := s.(Chinese1)
	//	ch1.niuYangGe()
	//case American1:
	//	//fmt.Println("美国人不会扭秧歌")
	//	am := s.(American1)
	//	am.disco()
	//}

	//var ch1 Chinese1 = s.(Chinese1)
	//fmt.Println(ch1)
	//s.niuYangGe()
}

func main() {
	//c := Chinese1{}
	//greet1(c)
	//a := American1{}
	//greet1(a)
	//var s string
	//s := "string"
	//k, v := s.(string)

}
