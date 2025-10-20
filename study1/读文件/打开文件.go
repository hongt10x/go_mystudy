package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println()
	//打开文件
	file, err := os.Open("D:\\go_study\\src\\calc\\读文件\\demo.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("文件=%v", file)
	fmt.Printf("文件=%s", file)
	defer file.Close()

}
