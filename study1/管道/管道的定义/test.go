package main

import (
	"fmt"
	"time"
)

/*
func main() {
	ch := make(chan int, 10)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i * i
			fmt.Println("写入", i*i)
			time.Sleep(1 * time.Second)

		}
		// The close built-in function closes a channel, which must be either
		// bidirectional or send-only. It should be executed only by the sender,
		// never the receiver, and has the effect of shutting down the channel after
		// the last sent value is received. After the last value has been received
		// from a closed channel c, any receive from c will succeed without
		// blocking, returning the zero value for the channel element. The form
		//    x, ok := <-c
		// will also set ok to false for a closed channel.
		close(ch)
		fmt.Println("管道已关闭，写入协程结束")
	}()
	go func() {
		for x := range ch {
			fmt.Println("读出", x)
		}
		fmt.Println("读取协程结束")
	}()

	time.Sleep(6 * time.Second)
	fmt.Println("GAME OVER")
}

*/

/*
func main() {
	queue := make(chan int, 1)
	//for data := range queue {
	//	fmt.Printf(strconv.Itoa(data) + " ")
	//}
	fmt.Println(len(queue), cap(queue))
	go func() {
		for {
			data := <-queue      //读取
			fmt.Print(data, " ") //0 1 2 3 4 5 6 7 8 9
		}
	}()

	for i := 0; i < 10; i++ {
		queue <- i //写入
		time.Sleep(time.Second)

	}
	time.Sleep(time.Second)
}


*/

func main() {
	// 创建一个无缓冲的通道
	queue := make(chan int)

	// 启动一个goroutine来发送数据到通道
	go func() {
		for i := 0; i < 10; i++ {
			queue <- i                         // 发送数据到通道
			time.Sleep(100 * time.Millisecond) // 模拟耗时操作
		}
		close(queue) // 发送完所有数据后关闭通道
	}()
	/*
		// 从通道中读取数据，直到通道被关闭
		for data := range queue {
			fmt.Print(data, " ") // 打印接收到的数据
		}
		fmt.Println()
		for i := 1; i < 10; i++ {

			fmt.Println("i:", i, <-queue)
		}
		fmt.Println() // 打印换行符以美化输出

	*/
	fmt.Println(len(queue))
	for i := 0; i < 10; i++ {
		fmt.Println(<-queue)
		fmt.Println("len: ", len(queue))

	}
}
