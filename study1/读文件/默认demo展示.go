package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	//读取的文件内容较少较小可用
	content, err := ioutil.ReadFile("D:\\go_study\\src\\calc\\读文件\\demo.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("File contents: %v", content)
	fmt.Printf("File contents: %s", content)

}
