package main

import (
	"fmt"
	"sync"
)

func main() {
	var dict sync.Map
	dict.Store(123, "123")
	dict.Store("abc", 123)
	t1 := test{name: "wang"}
	dict.Store(t1, "test")
	dict.Store("test1", t1)
	dict.Store(true, "true")
	fmt.Println(dict.Load("test11"))

	dict.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return false // 遇到false则终止遍历
	})
	fmt.Println("-----------------")
	dict.Range(func(key, value any) bool {
		fmt.Println(key, value)
		return true
	})
}

type test struct {
	name string
}
