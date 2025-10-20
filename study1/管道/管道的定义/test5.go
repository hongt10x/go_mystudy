package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3)
	var p1 []int16 = make([]int16, 3)
	fmt.Println("初始：", p1[0])
	go test1(ch)
	go test2(p1)

	time.Sleep(time.Second)

	fmt.Println("修改：", p1[0])
	r1 := <-ch
	fmt.Println(r1)
	fmt.Println("main.......")

}

func test1(ch chan int) {
	fmt.Println("test1 ......")
	ch <- 1000

}
func test2(p1 []int16) {
	fmt.Println("test2 ......")
	p1[0] = 99
	fmt.Println("修改中：", p1[0])
}
