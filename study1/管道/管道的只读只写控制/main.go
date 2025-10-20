package main

import "fmt"

func main() {
	//默认可读可写的管道
	//intChan1 := make(chan int)

	//声明只写管道
	var intChan2 chan<- string
	intChan2 = make(chan string, 3)
	intChan2 <- "1"
	intChan2 <- "1"
	intChan2 <- "1"

	//invalid operation: cannot receive from send-only channel intChan2 (variable of type chan<- int)
	//fmt.Println(<-intChan2) //只读

	//声明只读管道 [chan receive (nil chan)]:如果管道为空，则读取失败报错
	var intChan3 <-chan int
	fmt.Println(intChan3)
	if intChan3 != nil {
		fmt.Println(<-intChan3)
	}
	//invalid operation: cannot send to receive-only channel intChan3 (variable of type <-chan int)
	//intChan3 <- 100

}
