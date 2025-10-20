package main

import "fmt"

type Sing interface {
	singHi()
}

type Person1 struct {
}

func (p Person1) singHi1() {
	fmt.Println("有人在深夜唱歌...")
}

func (p Person1) singHi() {
	fmt.Println("开始接口测试...")
}

func sing(s Sing) {
	fmt.Printf("Sing: %T\n", s)
	s.singHi()
}

func sing1(p Person1) {
	fmt.Printf("Person1: %T\n", p)
	p.singHi()
}
func main() {
	p := Person1{}
	p.singHi1()

	p.singHi()

	sing(p)

	sing1(p)
}
