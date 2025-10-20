package main

import "fmt"

//var intChan1 chan int
//var stringChan chan string

func main() {
	//intChan1 = make(chan int, 3)
	//intChan1 <- 1
	//intChan1 <- 2
	//fmt.Printf("值%v，地址%v\n", <-intChan1, &intChan1)
	//fmt.Printf("值%v，地址%v\n", <-intChan1, &intChan1)
	////fmt.Printf("值%v，地址%v\n", <-intChan1, &intChan1)
	//stringChan = make(chan string, 3)
	//stringChan <- "wang"
	//stringChan <- "hongtao"
	//fmt.Printf("值%v，地址%v\n", <-stringChan, &stringChan)
	//fmt.Printf("值%v，地址%v\n", <-stringChan, &stringChan)

	allChan := make(chan interface{}, 10)
	allChan <- dog{Name: "alex", Color: "黄色"}
	allChan <- dog{Name: "bill", Color: "黑色"}
	allChan <- 123
	allChan <- "goodbye"
	//dog1 := <-allChan
	//fmt.Printf("%v %T\n", dog1, dog1)
	//gooddog := dog1.(dog) //经过断言后方可调用结构体方法
	//fmt.Printf("%v\n", gooddog.Name)
	//gooddog1 := (<-allChan).(dog) //经过断言后方可调用，如不存在则报错
	//fmt.Printf("%v\n", gooddog1.Name)

	//close(allChan)
	//for val := range allChan { //必须关闭管道，关闭后不能再写，才不会报错
	//	fmt.Println(val)
	//}

	//close(allChan)
	//fmt.Println(len(allChan))
	//for i := 0; i < len(allChan); i++ { //管道的长度len(allChan)因为取值而变化,需要注意
	//	fmt.Println(<-allChan)
	//}

	close(allChan)
	for {
		msg, ok := <-allChan
		if !ok {
			break
		}
		fmt.Println(msg, ok)
	}
}

type dog struct {
	Name  string
	Color string
}
