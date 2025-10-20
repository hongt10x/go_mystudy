package main

import (
	"fmt"
	"sync"
)

var totalNum int
var wg1 sync.WaitGroup

// 加入互斥锁
var lock sync.Mutex

// 循环1000次的结果还正常，到10000次就异常了
func add() {
	defer wg1.Done()
	for i := 0; i < 10000; i++ {
		lock.Lock() //加锁
		totalNum += 1
		lock.Unlock() //解锁
	}
}

func sub() {
	defer wg1.Done()
	for i := 0; i < 10000; i++ {
		lock.Lock()
		totalNum -= 1
		lock.Unlock()
	}
}

func main() {
	//两个协程操作同一个变量，引起变量输出不符合预期
	wg1.Add(2)
	go add()
	go sub()
	wg1.Wait()
	fmt.Println("totalNum = ", totalNum)
}
