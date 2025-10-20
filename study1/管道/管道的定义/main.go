package main

import (
	"fmt"
	_ "os"
)

func main() {
	//定义一个管道
	var intChan chan int
	//通过make进行初始化,可以存放3个int类型的数据
	intChan = make(chan int, 3)
	fmt.Printf("intChan=%p\n", intChan)

	//向管道存入数据
	intChan <- 1
	num := 20
	intChan <- num
	intChan <- 500
	//intChan <- 5001
	fmt.Printf("管道的实际长度是：%v\n", len(intChan))
	fmt.Printf("管道的容量是：%v\n", cap(intChan))

	//num1 := <-intChan
	//num2 := <-intChan
	//num3 := <-intChan
	//fmt.Printf("num1=%d, num2=%d, num3=%d\n", num1, num2, num3)

	//for i := 0; i <= len(intChan); i++ {
	//	fmt.Println("****")
	//	num1 := <-intChan
	//	fmt.Printf("num1=%d\n", num1)
	//}
	//var arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//for k, v := range arr {
	//	fmt.Println(k, v)
	//}

	close(intChan)
	//intChan <- 1999

	num4 := <-intChan
	num5 := <-intChan
	num6 := <-intChan
	num7 := <-intChan
	num8 := <-intChan
	fmt.Printf("num4=%v,%p\n", num4, &num4)
	fmt.Printf("num5=%v,%p\n", num5, &num5)
	fmt.Printf("num6=%v,%p\n", num6, &num6)
	fmt.Printf("num7=%v,%p\n", num7, &num7)
	fmt.Printf("num8=%v,%p\n", num8, &num8)

	ch := make(chan int, 10)
	for i := 0; i < 5; i++ {
		ch <- i * i
	}
	fmt.Printf("ch=%v %v\n", len(ch), cap(ch))
	close(ch)

	//for x := range ch {
	//	fmt.Println(x)
	//}
	for j := 0; j < len(ch)+5; j++ {
		fmt.Println(<-ch)
	}

}

/*
num4=1,0xc00009e090
num5=20,0xc00009e098
num6=500,0xc00009e0a0
*/
