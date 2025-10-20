package main

import (
	"fmt"
	"time"
)

//func main() {
//	var readOnlyChan <-chan int
//	var writeOnlyChan chan<- int
//	var ch chan int
//	ch = make(chan int, 3)
//	for i := 0; i < 3; i++ {
//		ch <- i
//	}
//	readOnlyChan = make(<-chan int, 2)
//	writeOnlyChan = make(chan<- int, 4)
//	fmt.Printf("readOnlyChan:%v\n", readOnlyChan)
//	fmt.Printf("writeOnlyChan:%v\n", writeOnlyChan)
//	fmt.Printf("ch:%v\n", ch)
//}

func main() {
	var ch chan int = make(chan int, 10)
	//for i := 0; i < 10; i++ {
	//	ch <- i
	//}
	//close(ch)
	//fmt.Println(<-ch)
	//<-ch
	//for v := range ch {
	//	fmt.Println(v)
	//}

	go writeChan(ch)
	go readChan(ch)
	time.Sleep(time.Second)
}

func readChan(ch <-chan int) { //从管道里出
	for v := range ch {
		fmt.Println(v)
	}
}

func writeChan(ch chan<- int) { //往管道里进
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}
