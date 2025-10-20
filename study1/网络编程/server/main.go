package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	for { //创建切片，将读取数据存放里面
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			return
		}
		fmt.Println(string(buf[:n]))
	}
}

func main() { //先启动服务端
	fmt.Println("服务器已启动...")

	//启动监听
	listen, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("监听异常： ", err)
		return
	}

	//监听成功，等待客户端连接
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("客户端连接失败：", err)
		} else {
			defer conn.Close()
			fmt.Printf("服务端信息：%v,接收到客户端信息：%v \n", conn.LocalAddr(), conn.RemoteAddr())
		}

		//准备协程，用来处理多方客户端的请求数据
		go process(conn)

	}

}
