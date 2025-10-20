package main

import "fmt"

func main() {
	var arr [5]int = [5]int{1, 2, 3, 4, 5}

	var slice = arr[2:5]
	fmt.Println(slice)

	var slice1 = make([]int, 3, 5)
	fmt.Println(slice1)

	for i, v := range slice1 {
		fmt.Println(i, v)
	}
	fmt.Println("******************")
	var slice2 []int
	//slice2 = make([]int, 3, 5)
	if slice2 == nil {
		fmt.Println("nil slice")
	} else {
		fmt.Println(slice2)
	}
}
