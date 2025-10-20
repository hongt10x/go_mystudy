package main

import "fmt"

func main() {
	fmt.Println("hello world1")
	fmt.Println("hello world2")
	goto lable1
	fmt.Println("hello world3")
	fmt.Println("hello world4")
	fmt.Println("hello world5")
	fmt.Println("hello world6")
lable1:
	fmt.Println("hello world7")
	fmt.Println("hello world8")
}

/*
hello world1
hello world2
hello world7
hello world8
*/
