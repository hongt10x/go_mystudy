package main

import "fmt"

/*
	变量存储：
		等号 左边的变量，代表 变量所指向的内存空间。	（写）
		等号 右边的变量，代表 变量内存空间存储的数据值。	（读）

	指针的函数传参（传引用）：
		传地址（引用）：将实参的地址值作为函数参数传递。
		传值（数据）：将实参的值拷贝一份给形参。
		传引用：在A栈帧内部，修改B栈帧中的变量值。

*/

//func main() {
//	var num int = 100
//	fmt.Printf("%v %p %v\n", num, &num, *&num)
//	var x int
//	x = num
//	fmt.Printf("%v %p %v\n", x, &x, *&x)
//	var y *int
//	y = &num
//	fmt.Printf("%v %p %v\n", y, &y, *y)
//
//	//	var a *int = &num
//	//	var b *int = new(int)
//	//	fmt.Println(a, b, *a, *b)
//	//	a, b = b, a
//	//	fmt.Println(a, b, *a, *b)
//}

func main() {
	var n int = 100
	fmt.Printf("%v %p\n", n, &n)
	var p *int = &n
	fmt.Printf("%v %v\n", p, *(&n))
	*p = 200
	fmt.Println(*p, *&n, n)

}
