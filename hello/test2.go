package main

import "fmt"

func main() {
	fmt.Println("test111")
	l1 := make(map[int]int)
	for i := 0; i < 10; i++ {
		l1[i] = i + 1
	}
	fmt.Println(l1)

	for k, v := range l1 {
		fmt.Println(k, v)
	}
}
