package main

import "fmt"

func main() {
	var arr1 = [3]int{1, 3, 4}
	fmt.Println(arr1)

	var arr2 = [...]int{1, 3, 4, 5, 6}
	fmt.Println(arr2)
	//fmt.Printf("%p\n", &arr2)
	//fmt.Printf("%p\n", &arr2[0])
	//fmt.Printf("%p\n", &arr2[1])
	//fmt.Printf("%T\n", &arr2)
	//fmt.Printf("%T\n", &arr2[0])
	test11(arr1)
	fmt.Println(arr1)
	test12(&arr1)
	fmt.Println(arr1)
}

// 值传递
func test11(arr [3]int) {
	arr[0] = 999
	fmt.Println(arr)
}

// 地址传递
func test12(arr *[3]int) {
	arr[0] = 888
	fmt.Println(*arr)
}

/*

func calc() {
	//一个值被改动，其他引用也将随着变化
	s1 := []int{1, 2, 3, 4, 5}
	fmt.Println(s1)
	//[...]该声明为已数组的长度，由编译器计算，根据实际的赋值设定长度，一旦确定，不可更改
	s22 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(s22)
	s22[4] = 100

	//s2 := s1[:3]
	//fmt.Println(s2)
	//s2[0] = 100
	//fmt.Println(s2)
	//fmt.Println(s1)
	//s3 := s2[1:]
	//fmt.Println(s3)
	//s3[0] = 999
	//fmt.Println(s3)
	//fmt.Println(s1)
	////fmt.Println(s2)

	s11 := []int{1, 2, 3, 4, 5}
	fmt.Println(s11)
	s12 := append(s11, 7, 7, 8, 9)
	fmt.Println(s12)
	s12[0] = 100
	s12[1] = 200
	fmt.Println(s12)
	fmt.Println(s11)
	//主意直接追加切片需要再后边加...,三个点必写不可省略
	s13 := append(s11, s12...)
	fmt.Println(s13)
	s144 := append(s11, s13...)
	fmt.Println(s144)

	var s14 = make([]int, 10, 99)
	copy(s14, s13)
	fmt.Println("--------------")
	fmt.Println(s13)
	fmt.Println("s144:", s144)
	fmt.Printf("s144: %p\n", s144)
	var s145 []int //初始切片是空 []
	fmt.Println(s145)
	s145 = append(s145, 111)
	copy(s145, s144)
	fmt.Println("s145", s145)
	fmt.Printf("s145: %p\n", s145)

	var s156 [5]int // 初始数组是[0 0 0 0 0]，而非空
	fmt.Println(s156)

	var s157 []int //[]
	fmt.Println(s157)

}


*/
