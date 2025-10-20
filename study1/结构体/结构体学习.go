package main

import (
	"fmt"
	"reflect"
)

type Teacher struct {
	Name   string
	Age    int
	School string
}

func main() {
	fmt.Println()
	/*
		//方式1：
		//一个对象有很多属性情况，用变量管理太分散
		var t1 Teacher
		//有默认值，空或0
		fmt.Printf("t1: %T, %v\n", t1, t1)
		t1.Name = "马学兵"
		t1.Age = 45
		t1.School = "清华大学"
		fmt.Println(t1)
		fmt.Println(t1.School)
	*/
	//方式2:
	var t2 = Teacher{}
	fmt.Printf("t2: %T, %v\n", t2, t2)
	t2.Name = "马学兵"
	fmt.Println(t2)
	t2.Age = 45
	fmt.Println(t2)
	fmt.Println("School:", t2.School)
	t2.School = "清华大学"
	fmt.Println(t2)
	fmt.Println("==================")

	rv := reflect.ValueOf(t2)
	fmt.Println(rv.Kind(), "****", rv)

	rv2 := reflect.TypeOf(t2)
	fmt.Println(rv2)
	fmt.Println("=======000000===========")
	//方式3：
	//var t3 Teacher = Teacher{"赵姗姗", 32, "黑龙江大学"}
	//fmt.Println(t3)
	//fmt.Println(t3.Name)

	//方式4：指针类型
	var t4 *Teacher = new(Teacher)
	fmt.Println(*t4)
	fmt.Printf("%T, %p,%v\n", t4, t4, t4)
	//标准写法
	(*t4).Age = 44
	fmt.Println(*t4)
	fmt.Println((*t4).Age)

	//go为了程序员的书写习惯，给做了一层转化，原理同上
	t4.Name = "张三"
	fmt.Println(*t4)

	//方式5：
	var t5 *Teacher = &Teacher{}
	(*t5).Age = 44
	t5.Name = "兔子"
	fmt.Println(*t5)

	var t6 = new(Teacher)
	t6.School = "武汉大学"
	(*t6).Age = 444
	t6.Name = "wht"
	fmt.Println(*t6)

}
