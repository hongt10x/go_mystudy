package main

import "fmt"

func main() {
	a, b := test1(1)
	fmt.Println(a, b)
}

func test1(i int) (bool, int) {
	fmt.Println(i)
	return true, 1
}
