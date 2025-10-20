package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("D:\\go_study\\src\\calc\\读文件\\demo.txt")
	if err != nil {
		fmt.Printf("err: %T\n", err)
		fmt.Println("err:", err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')

		//fmt.Printf(str)
		fmt.Printf("str=%v", str)
		if err == nil {
			fmt.Printf("err=%T\n", err)
		}

		if err == io.EOF {
			fmt.Println()
			fmt.Printf("err=%v\n", err)
		}
		if err == io.EOF {
			break
		}
	}
	fmt.Println("end")

}

//func calc() {
//
//	content, err := ioutil.ReadFile("D:\\go_study\\src\\calc\\读文件\\demo.txt")
//	if err != nil {
//		panic(err)
//	} else {
//		fmt.Println(string(content))
//	}
//}
