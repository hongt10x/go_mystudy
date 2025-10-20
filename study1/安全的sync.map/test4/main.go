package main

import "fmt"

func counter(ch chan int) {
	for i := 1; i <= 5; i++ {
		fmt.Println("当前值", i)
		ch <- i
		fmt.Println("写入管道", i)
	}
	close(ch)
}

func main() {
	ch := make(chan int, 0)
	go counter(ch)
	for num := range ch {
		fmt.Println(num)
	}
}
