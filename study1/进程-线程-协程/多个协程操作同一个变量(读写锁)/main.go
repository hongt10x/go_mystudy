package main

import (
	"fmt"
	"sync"
	"time"
)

var wg_rw sync.WaitGroup
var lock_rw sync.RWMutex

func read1() {
	defer wg_rw.Done()
	lock_rw.RLock()
	fmt.Println("开始读取")
	time.Sleep(time.Second)
	fmt.Println("读取完成")
	lock_rw.RUnlock()
}

func write1() {
	defer wg_rw.Done()
	lock_rw.Lock()
	fmt.Println("开始写入")
	time.Sleep(time.Second * 3)
	fmt.Println("写入完成")
	lock_rw.Unlock()
}

func main() {

	wg_rw.Add(20)

	for i := 0; i < 16; i++ {
		go read1()
	}
	for j := 0; j < 4; j++ {
		go write1()
	}
	wg_rw.Wait()
	fmt.Println("Done...")
}
