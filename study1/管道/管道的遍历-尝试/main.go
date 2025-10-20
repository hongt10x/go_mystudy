package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建一个整型管道
	ch := make(chan int, 2)

	// 启动一个goroutine来发送数据到管道
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i                 // 发送数据到管道
			time.Sleep(time.Second) // 模拟耗时操作
		}
		close(ch) // 发送完毕后关闭管道
	}()

	// 使用for循环从管道中读取数据
	for {
		// 从管道中接收数据
		val, ok := <-ch
		if !ok {
			// 如果管道已关闭且没有数据可读，则退出循环
			break
		}
		fmt.Println(val) // 打印接收到的数据
	}

	fmt.Println("数据读取完毕")
}
