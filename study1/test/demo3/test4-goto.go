package main

import "fmt"

func main() {
	fmt.Println("start")
	goto flag
	fmt.Println("a")
flag:
	fmt.Println("b")
}
