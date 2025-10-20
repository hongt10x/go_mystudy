package main

import "fmt"

/*
func main() {
	//当管道阻塞时，有两种方法可以解决：1、取值，读取管道中的数据进行释放；2、增大管道的容量；
	ch := make(chan int, 1)
	ch <- 1

	fmt.Println(<-ch)
	ch <- 2
	fmt.Println(<-ch)
	//close(ch)
	for val := range ch {
		fmt.Println(val)
	}

}


*/

func main() {
	// 创建一个整型通道
	ch := make(chan int, 1)

	// 尝试将0通过通道发送
	ch <- 110

	data, ok := <-ch
	fmt.Println(data, ok)

	//ch <- 100
	//data1, ok1 := <-ch
	//
	//fmt.Println(data1, ok1)
}
