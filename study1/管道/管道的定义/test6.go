package main

import (
	"fmt"
	"time"
)

func main() {
	GOTIME := "2006-01-02 15:04:05"
	fmt.Println(time.Now().Format(GOTIME))
}
