package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
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

	//通过客户端发送数据到服务端，然后退出
	reader := bufio.NewReader(os.Stdin)
	str, err2 := reader.ReadString('\n')
	if err2 != nil {
		fmt.Println("终端输入失败：", err)
	}
	//将str发送给服务端
	n, err1 := conn.Write([]byte(str))
	if err1 != nil {
		fmt.Println("连接失败：", err)
	}
	fmt.Println("客户端发送数据为：", n)

}
