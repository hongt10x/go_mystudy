package main

import "fmt"

func main() {
	fmt.Println("i am main")
	f1()
	f2()
}

/*
需要将该包裹下所有文件的package都命名为main才可以运行，否则报错。

问题1：为什么需要将该包裹下所有文件的package都命名为main才可以运行？

而且好像只能在文件夹那里右键运行才行，如果直接在代码编辑区右键运行会报错:

问题2：为什么直接在代码编辑区右键运行会报错？
*/
