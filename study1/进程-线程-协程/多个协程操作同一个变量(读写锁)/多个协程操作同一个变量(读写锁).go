package main

import (
	"fmt"
	"sync"
	"time"
)

// var totalNum int
var wg1 sync.WaitGroup

// 加入互斥锁
var lock sync.RWMutex

// 循环1000次的结果还正常，到10000次就异常了
func read() {
	defer wg1.Done()
	lock.RLock()
	fmt.Println("开始读取数据")
	time.Sleep(time.Second)
	fmt.Println("读取数据成功")
	lock.RUnlock() //解锁
}

func write() {
	defer wg1.Done()
	lock.Lock()
	fmt.Println("开始写入数据...")
	time.Sleep(time.Second * 5)
	fmt.Println("写入数据成功...")
	lock.Unlock() //解锁
}

func main() {
	//两个协程操作同一个变量，引起变量输出不符合预期
	defer wg1.Wait()
	wg1.Add(11)
	for i := 0; i < 10; i++ {
		go read()
	}

	go write()

}
