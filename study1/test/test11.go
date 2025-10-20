package main

import "fmt"

/*
func calc() {

	var str1 string = "hello world 中国！"
	fmt.Println(str1)

	fmt.Println(time.Now())
	fmt.Println(time.DateTime)
	fmt.Println(time.Now().Year())
	fmt.Println(time.Now().Month())
	fmt.Println(time.Now().Day())
	//格式化的数字必须是固定的，必须要这样写2006-01-02 15:04:05，这是个模板写法不可更改，否则识别为其他
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println(time.Now().Format("2006 15:04:05"))
	fmt.Println(time.Now().Format("2006 15:04:51"))
}


*/

func main() {
	//t1 := make([]int, 10)
	//for i := 0; i < 10; i++ {
	//	t1[i] = i
	//}
	//t2 := make([]int, 10)
	//fmt.Println(t1, t2)
	//copy(t2, t1)
	//fmt.Println(t1, t2)

	t33 := make([][]int, 10)
	t33[0] = append(t33[0], 100)
	fmt.Println(t33)
	t44 := [][]int{{1, 2}, {3, 5}}
	fmt.Println(t44)
	t44[0] = append(t44[0], 100)
	fmt.Println(t44)

	//slice1 := []int{1, 2, 3, 4, 5}
	//slice2 := []int{5, 4, 3}
	//fmt.Println(slice1, slice2)
	//copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	//fmt.Println(slice1, slice2)
	//copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置
	//fmt.Println(slice1, slice2)
}
