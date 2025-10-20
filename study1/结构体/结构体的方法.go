package main

import "fmt"

type Person1 struct {
	Name string
	Age  int
}

// 给Person绑定方法
func (p Person1) test() {
	fmt.Println(p.Name)
	fmt.Println(p.Age)
}

func (Person1) test1() {
	fmt.Println(Person1.test1)
	fmt.Println(Person1.test34)

}

func test33(a, b int) int {
	return a + b
}
func (p Person1) test34(a, b int) int {
	return a + b
}

func test35(a, b int) (int, int) {
	return b, a
}
func (p Person1) test36(a, b int) (int, int) {
	return b, a
}

func main() {
	fmt.Println()
	var p Person1
	p.Name = "丽丽"
	p.Age = 20
	p.test()
	fmt.Println(p.test1)
	p.test1()
	fmt.Println(p.test36(10, 5))
	fmt.Println(test33(3, 4))
	fmt.Println(p.test34(4, 5))
	a1, b1 := test35(2, 5)
	fmt.Println(a1, b1)

	//for i := 0; i < 10; i++ {
	//	p.test1()
	//}

	fmt.Println("------------------")

	var x Person1
	var y = new(Person1)
	var z = &Person1{}
	fmt.Printf("%p,%v\n", &x, x)
	fmt.Printf("%p,%v\n", y, *y)
	fmt.Printf("%p,%v\n", z, *z)

	fmt.Println(x.Name)
	x.Name = "good"
	fmt.Println(x)
	x.Age = 100
	fmt.Println(x)
	var tea1 = Person1{"alex", 30}
	var tea2 = Person1{Name: "alex", Age: 30}
	fmt.Println(tea1, tea2)
	fmt.Println(tea2.Name, tea2.Age)
}
