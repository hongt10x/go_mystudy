package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg3 sync.WaitGroup

func main() {
	intChan3 := make(chan int, 50)
	exitChan := make(chan bool, 1)
	//fmt.Println(<-exitChan)
	//wg3.Add(1)
	go write(intChan3)
	go read(intChan3, exitChan)
	//wg3.Wait()
	fmt.Println("main end...1")
	if <-exitChan {
		return
	}
	//fmt.Println(<-exitChan)
	fmt.Println("main end...2")
}

func write(intChan chan int) {
	//defer wg3.Done()
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= 50; i++ {
		temp := rand.Intn(100) + 10
		fmt.Printf("第%v次写数，数据内容为：%v\n", i, temp)
		intChan <- temp
		//time.Sleep(time.Second)
	}
	close(intChan)
}

func read(intChan chan int, exitChan chan bool) {
	var count int
	for {
		tmp, ok := <-intChan
		count++
		if !ok {
			break
		}
		fmt.Printf("第%v次读取，数据内容为：%v\n", count, tmp)
	}
	exitChan <- true //读取完毕
	//close(intChan)
}
