package main

import (
	"fmt"
	"reflect"
)

func testRef(i interface{}) {
	//typeof
	reType := reflect.TypeOf(i)
	fmt.Println(reType)
	//valueof
	val := reflect.ValueOf(i)
	fmt.Println(val)

	//类别 > 类型
	refKind := reType.Kind()
	fmt.Printf("%T,%v\n", refKind, refKind)
	refKind1 := val.Kind()
	fmt.Printf("%T,%v\n", refKind1, refKind1)

	i1 := val.Interface()
	if v1, ok := i1.(int); ok {
		fmt.Println(v1)
	} else {
		fmt.Println("not int")
	}

}

type Student struct {
	name string
	age  int64
}

func main() {
	//stu := Student{
	//	name: "旺旺",
	//	age:  20,
	//}
	//testRef(stu)
	i := 100000000011111111
	testRef(i)
}
