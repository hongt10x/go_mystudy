package main

import "fmt"

func main() {
	var arr [2][3]int16
	fmt.Println(arr)
	fmt.Printf("%p\n", &arr)
	fmt.Printf("%p\n", &arr[0])
	fmt.Printf("%p\n", &arr[0][0])
	fmt.Printf("%p\n", &arr[0][1])
	fmt.Printf("%p\n", &arr[1])
	fmt.Printf("%p\n", &arr[1][0])
	fmt.Printf("%p\n", &arr[0][1])

	//var arr2 [2][3]int16 = [2][3]int16{{1, 4, 7}, {2, 5, 8}}
	//fmt.Println(arr2)
	//
	var arr3 = [...][3]int16{{1, 2, 3}, {8, 9}, {3}}
	fmt.Println(arr3)
	//当你使用切片时，内层数组（在这个例子中是切片）的长度可以不同。例如，[][]int16{{1, 2}, {8, 9, 0, 1}} 是有效的，但如果你使用二维数组，则所有内层数组的长度必须相同。
	//切片在 Go 中是非常强大和灵活的数据结构，它们允许你以类似于动态数组的方式工作，同时保持高性能。
	var arr4 = [][]int16{{1, 2, 3}, {8, 9}}
	fmt.Println(arr4)

	//循环读取方式1：
	for i := 0; i < len(arr4); i++ {
		for j := 0; j < len(arr4[i]); j++ {
			fmt.Print(arr4[i][j], "\t")
		}
		fmt.Println()
	}

	//循环读取方式2：
	fmt.Println("---------------------")
	for key, value := range arr4 {
		for k, v := range value {
			fmt.Printf("arr[%v][%v]=%v	", key, k, v)
		}
		fmt.Println()
	}
}
