package main

import "fmt"

func test1(a, b *int) {
	fmt.Println(*a, *b)
	fmt.Println(a, b)
	*a, *b = *b, *a
}

func test2(a, b *int) int {
	fmt.Println(*a, *b)
	fmt.Println(a, b)
	fmt.Printf("%v %v\n", a, b)
	a, b = b, a
	return 0
}

func main() {
	var x, y = 10, 20
	fmt.Println(x, y)
	fmt.Println(&x, &y)
	test1(&x, &y)
	fmt.Println(&x, &y)
	fmt.Println(x, y)

	var z = 100
	fmt.Println(z)

	test2(&z, &x)
	fmt.Println(&x, &y)
	fmt.Println(x, y)

}
