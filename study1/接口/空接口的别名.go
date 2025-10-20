package main

import "fmt"

func main() {
	test(1)
	test("abc")
	mt := myTest{name: "alex"}
	test(mt)
	fmt.Println("++++++++++++++++++++")
	test1(123)
	test1("abcdef")
	mt1 := myTest{name: "echo"}
	test1(mt1)
}

// any
func test(actual any) {
	fmt.Println(actual)
}

// 空接口
func test1(actual interface{}) {
	fmt.Println(actual)
}

type myTest struct {
	name string
}
