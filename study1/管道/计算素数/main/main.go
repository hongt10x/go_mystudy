package main

import "fmt"

var intChan chan int

func iknow() {
	fmt.Println("iknow")
	if err := recover(); err != nil {
		fmt.Println(err)
	}
}

func check(a int, b int) {
	//defer func() {
	//	if err := recover(); err != nil {
	//		fmt.Println(err)
	//	}
	//}()
	intChan := make(chan int, 1)

	intChan <- 1
	close(intChan)
	c := a / b
	fmt.Println(c)
}
func main() {
	defer iknow()
	//check(1, 0)
	k := 3 / 0
	fmt.Println(k)

	fmt.Println("bye")

}
