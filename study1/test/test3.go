package main

import "fmt"

func main() {
	var input int
	fmt.Println("请输入数字：")
	fmt.Scanln(&input)
	switch input {
	case 4:
		{
			fmt.Println("4")

		}
		fallthrough
	case 3:
		{
			fmt.Println("3")
		}
	case 2:
		{
			fmt.Println("2")
		}
	default:
		{
			fmt.Println("default")
		}
	}
}

//func calc() {
//	var count int = 8
//	//if count < 10 {
//	//	fmt.Println("库存不足")
//	//} else {
//	//	fmt.Println("库存很多")
//	//}
//
//	switch count / 2 {
//	case 4, 3:
//		println("4,3")
//		//仅穿透下面一层
//		//fallthrough
//
//	case 2:
//		println("2")
//	case 1:
//		println("1")
//	default:
//		println("无效")
//
//	}
//}
