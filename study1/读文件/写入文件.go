package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//perm 只在linux/unix有效
	file, err := os.OpenFile("D:/go_study/src/calc/读文件/管道的定义.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("你好 goland\n")
	}

	writer.Flush()
	s := os.FileMode(0666).String()
	fmt.Println(s)
}

/*

func OpenFile(name string, flag int, perm FileMode) (*File, error)
该方法第一个参数为文件路径,第二个参数控制文件的打开方式,第三个参数控制文件模式
可用的打开方式有
// Flags to OpenFile wrapping those of the underlying system. Not all
// flags may be implemented on a given system.
const (
    // Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
    // 只读模式
    O_RDONLY int = syscall.O_RDONLY // open the file read-only.
    // 只写模式
    O_WRONLY int = syscall.O_WRONLY // open the file write-only.
    // 可读可写
    O_RDWR   int = syscall.O_RDWR   // open the file read-write.
    // The remaining values may be or'ed in to control behavior.
    // 追加内容
    O_APPEND int = syscall.O_APPEND // append data to the file when writing.
    // 创建文件，如果文件不存在
    O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
    // 与创建文件一同使用，文件必须存在
    O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must exist.
    // 打开一个同步的文件流
    O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
    // 如果可能，打开时缩短文件
    O_TRUNC  int = syscall.O_TRUNC  // truncate regular writable file when opened.
)
*/
