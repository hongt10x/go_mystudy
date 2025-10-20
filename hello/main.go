package main

import (
	"fmt"

	"hello/users"
)

//
//import "fmt"
//
//func main() {
//	fmt.Println("main")
//	test111()
//}

/*
https://www.cnblogs.com/huzhengyu/p/14584747.html
执行：go build .\main.go .\test.go ，执行编译后的 main.exe ，正常运行

执行：go run main.go test.go，正常运行

所以，第一种运行方式，编译器并不把两个文件当一个包运行。

将 hello文件夹，设置成工程目录，也可以运行

*/

func main() {
	fmt.Println("main")
	users.Usering()
	test111()
}
