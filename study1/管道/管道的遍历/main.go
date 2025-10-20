package main

import "fmt"

func main() {
	var intChan chan int
	intChan = make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan <- i + 10
	}
	close(intChan) //遍历之前一定要把管道关闭，代表管道写入完成

	//只能用for...range进行轮询读取
	for v := range intChan {

		fmt.Println(v)
	}
	//defer close(intChan)
	//close(intChan) //这个位置代表管道还没关闭，上面的循环会一直读取数据而出错fatal error: all goroutines are asleep - deadlock!
}
