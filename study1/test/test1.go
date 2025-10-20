package main

import (
	//d "demo1" // 对包进行别名，别名前是demo1.   别名后是 d.
	"fmt"
	//. "demo1"  //引用包时，取消包.  即  取消 demo1.   直接引用函数或变量即可
	_ "fmt" //屏蔽不再引用的包用_
)

var x int = 100
var y string = "string"

// k := fmt.Sprint(x)
const apple = "good apple"
const ORANGE = "good orange" //常量用大写

var b bool = true

// 声明一组变量
//var (
//	name   = "alex"
//	age    = 29
//	gender = "male"
//)

func main() {
	fmt.Println(apple)
	fmt.Println(ORANGE)
	//fmt.Println(d.MYNAME)
	//fmt.Println(d.myage)
	//fmt.Println(demo1.MYNAME)
	//fmt.Println("HELLO, Go !", apple)
	//fmt.Println("HELLO, Go !", y)
	//var z = "x=%d---y=%s"
	//var j = fmt.Sprintf(z, x, y)
	//fmt.Println(j)
	//k := fmt.Sprint(x)
	//fmt.Println(k)
	//fmt.Println(b)

	//var num int
	//fmt.Println(num)
	//fmt.Println(name)
	//fmt.Println(age)
	//fmt.Println(gender)
	//
	//var n2 float32 = 100
	//fmt.Println(n2)
	//var n3 float64 = float64(n2)
	//fmt.Println(n3)
	//fmt.Printf("%T", n3)

	//fmt.Println(demo1.UserName)

}
