package main

import (
	"fmt"
	"time"
)

func test() {
	for i := 1; i <= 100; i++ {
		fmt.Printf("hello golang + %v\n", i)
		time.Sleep(time.Second)
	}
}

// 主死从随 -- 主线程结束后，协程也将强制结束
func main() { //主线程
	go test() //开启一个协程
	for i := 1; i <= 10; i++ {
		fmt.Printf("hello msb + %v\n", i)
		time.Sleep(time.Second)
	}

}

/*输出
hello msb + 1
hello golang + 1
hello golang + 2
hello msb + 2
hello msb + 3
hello golang + 3
hello msb + 4
hello golang + 4
hello golang + 5
hello msb + 5
hello golang + 6
hello msb + 6
hello msb + 7
hello golang + 7
hello golang + 8
hello msb + 8
hello msb + 9
hello golang + 9
hello msb + 10
hello golang + 10
*/
