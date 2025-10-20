package main

import (
	"fmt"
	"reflect"
)

// 当你使用interface{}类型（也称为空接口）作为函数参数时，你可以传递任何类型的值给这个函数，因为空接口不定义任何方法，因此它满足了所有类型的“接口”。在你给出的例子中，testRef函数接收一个interface{}类型的参数i，这意味着它可以接收任何类型的值。
func testRef(i interface{}) {
	//typeof
	reType := reflect.TypeOf(i)
	fmt.Println("reType:", reType)
	fmt.Printf("%T,%v\n", reType, reType)
	//valueof
	val := reflect.ValueOf(i)
	fmt.Printf("%T,%v\n", val, val)
	fmt.Println("=================")
	fmt.Println(val.Kind())
	fmt.Printf("%T,%v\n", val.Kind(), val.Kind())
	fmt.Println(reflect.Struct)

	num := 100 + val.Int()
	fmt.Println("求和：", num)
	num2 := val.Interface()
	num3, ok := num2.(int)
	fmt.Println(num3)
	fmt.Println(ok)
}

func main() {
	num := 400
	testRef(num)
	fmt.Println(num)
}
