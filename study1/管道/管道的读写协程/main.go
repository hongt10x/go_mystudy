package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func writeData(intChan chan int) {
	defer wg.Done()
	for i := 1; i <= 50; i++ {
		intChan <- i
		fmt.Println("writeData: ", i)
		time.Sleep(time.Second)
	}

	//关闭管道
	close(intChan)
}

func readData(intChan chan int) {
	defer wg.Done()
	for v := range intChan {
		fmt.Println("readData: ", v)
	}
}

func main() {
	intChan := make(chan int, 50)
	wg.Add(2)
	//开启写入管道的协程
	go writeData(intChan)
	//开启读取的管道协程
	go readData(intChan)
	wg.Wait()
}

/*
...
readData:  48
writeData:  49
readData:  49
writeData:  50
readData:  50
*/
