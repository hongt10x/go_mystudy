package main

import "fmt"

func main() {
	a := 1
	res := a << 4
	fmt.Println(a, res) // 16
	b := 8
	res1 := b >> 2
	fmt.Println(b, res1) // 2
	c := a < 11 && b > 1
	fmt.Println(c)
	d := a < 1 || b > 8
	fmt.Println(d)
}
