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

	v := val.Interface()

	ret, ok := v.(Student)
	if ok {
		fmt.Println(ret)
	}

	//获取字段的数量
	n1 := val.NumField()
	fmt.Println(n1)
	//获取每个字段的内容
	for i := 0; i < n1; i++ {
		fmt.Printf("获取第%d个字段%v 值:%v\n", i+1, val.Field(i), val.Field(i))
	}
	//获取字段的名字
	fmt.Println(val.FieldByName("name"))
}

type Student struct {
	name string
	age  int
}

func main() {
	stu := Student{
		name: "旺旺",
		age:  20,
	}
	testRef(stu)
}
