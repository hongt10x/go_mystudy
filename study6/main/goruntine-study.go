package main

import (
	"fmt"
	"strconv"
	"time"
)

var ch chan int

func main() {
	ch = make(chan int)
	go func() {
		for i := 1; i < 100; i++ {
			if i == 3 {
				//runtime.Gosched() //主动出让执行
				<-ch //阻塞
			}
			fmt.Println("rountine 1: " + strconv.Itoa(i))
		}
	}()
	go func() {
		for i := 100; i < 200; i++ {
			fmt.Println("rountine 2: " + strconv.Itoa(i))
		}
		ch <- 1
	}()
	time.Sleep(time.Second * 5)
}
