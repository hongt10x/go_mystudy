package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("客户端启动...")

	//调用dial函数
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("客户端连接失败：", err)
		return
	}
	fmt.Println("连接成功：", conn.LocalAddr())

}
