package main

import (
	"fmt"
	"reflect"
)

func test(i any) {
	val1 := reflect.ValueOf(i)
	fmt.Printf("%T,%v\n", val1, val1)
	val1.Elem().SetInt(999)
	fmt.Printf("%T,%v\n", val1.Elem(), val1.Elem())

}

func test2(i interface{}) {
	//*i.(*int) = 1999
	if j, ok := i.(*int); ok {
		*j = 99900
		fmt.Println(*j)
	} else {
		fmt.Println(i)
	}

}

func main() {
	num := 40
	fmt.Printf("%v,%p\n", num, &num)
	test(&num)
	fmt.Println(num)
	//test2(&num)
	//fmt.Println(num)
}
