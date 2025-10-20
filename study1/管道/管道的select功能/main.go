package main

import (
	"fmt"
	"time"
)

func main() {
	//select处理多个管道问题
	//定义一个int管道
	intChan := make(chan int, 1)
	defer close(intChan)
	go func() {
		time.Sleep(time.Second * 5)
		intChan <- 10

	}()
	//定义一个string管道
	stringChan := make(chan string, 1)
	defer close(stringChan)
	go func() {
		time.Sleep(time.Second * 2)
		stringChan <- "golang"

	}()

	//读取数据本身就是阻塞
	//fmt.Println(<-intChan)
	for {
		select {
		case v := <-intChan:
			fmt.Println("intChan: ", v)
			return
		case v := <-stringChan:
			fmt.Println("stringChan: ", v)

			//default: //防止select被阻塞，默认执行default
			//	fmt.Println("default")
			//	break

		}
	}
}
