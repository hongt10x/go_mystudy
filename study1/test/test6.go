package main

import "fmt"

func main() {
	//for i := 0; i < 5; i++ {
	//	for j := 0; j < 5; j++ {
	//		fmt.Printf("%d , %d \n", i, j)
	//		if i == 2 && j == 2 {
	//			break
	//		}
	//	}
	//}
	//增加标签，可以标记主动跳出的循环
lable1:
	for i := 0; i < 5; i++ {
	lable2:
		for j := 0; j < 5; j++ {
			fmt.Printf("%d , %d \n", i, j)
			if i == 2 && j == 2 {
				break lable1
			}
			if i == 1 && j == 1 {
				break lable2
			}
		}
	}
	fmt.Println("-----ok------")

	//continue也可以加lable
lable:
	for i := 1; i < 5; i++ {

		for j := 1; j < 5; j++ {
			if i == 1 && j == 1 {
				continue lable
			}
			fmt.Printf("%d , %d \n", i, j)
		}
	}
	fmt.Println("-----ok------")
	for i := 0; i < 5; i++ {
		fmt.Printf("%d \n", i)
		if i == 1 {
			return
		}
	}
	fmt.Println("-----ok------")

}
