package main

import "fmt"

func main() {
	i := 0
	fmt.Println(&i)
	f(&i)
	fmt.Println(i)
}

func f(count *int) {
	fmt.Println("f -->", *count)
	(*count)++
	fmt.Println(count, *count)
}
