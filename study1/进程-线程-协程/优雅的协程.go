package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup //只定义无需复制
//func calc() {
//	fmt.Println("start calc...")
//	for i := 0; i < 10; i++ {
//		wg.Add(1) //计数器加一
//		go func(n int) {
//			defer wg.Done()
//			fmt.Println("hello coroutine ", n)
//			//wg.Done()  在函数执行结束后，计数器减一
//		}(i)
//	}
//	wg.Wait()
//	fmt.Println("stop calc...")
//
//}

func main() {
	fmt.Println("start calc...")
	//wg.Add(10) //计数器加一，需要注意要与协程的数量保持一致;这种方法需要计算协程数量，并不符合实际开发情况
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) { //开启一个协程
			defer wg.Done()
			fmt.Println("hello coroutine ", n)
			//wg.Done()  在函数执行结束后，计数器减一
			time.Sleep(1)
		}(i)
	}
	wg.Wait()
	fmt.Println("stop calc...")

}
