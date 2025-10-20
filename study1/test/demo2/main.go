package main

import "fmt"

func test1() {
	fmt.Println("hello world")
}
func main() {
	a := 100
	fmt.Println(a)
	b := a
	fmt.Println(b)
	test1()
	t := test1
	t()
}
