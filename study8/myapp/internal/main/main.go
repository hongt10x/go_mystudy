package main

import (
	"fmt"

	utils "study/study8/myapp/internal/myutils"
)

func main() {
	fmt.Println("main --> Hello World")
	utils.PrintHello() //package的名字与目录不同，相当于别名，需要留意
}
