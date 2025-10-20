package main

import "fmt"

func main() {
	//var intarr [5]int = [5]int{1, 2, 3, 4, 5}
	//fmt.Println(intarr)
	//s100 := intarr[2:5]
	//fmt.Println(s100, len(s100), cap(s100))
	//s100 = append(s100, 111, 222, 333, 44)
	//fmt.Println(s100, len(s100), cap(s100))
	k1 := make([]int, 3, 3)
	fmt.Println(k1)
	k1 = append(k1, 2)
	k1 = append(k1, 4)
	k1 = append(k1, 5)
	k1 = append(k1, 6)
	fmt.Println(k1)
}

/*
func calc() {

	//fmt.Println()
	//var intarr = [6]int{1, 2, 3, 4, 5}
	//fmt.Println(intarr)
	////定义一个动态变化的slice切片，左包右不包！！！！
	//var slice []int = intarr[3:5]
	//slice1 := intarr[3:]
	//fmt.Println(slice)
	//fmt.Println(slice1)
	//fmt.Println(len(slice1))
	//fmt.Println(cap(slice1))

	slice2 := make([]int, 4, 7)
	var slice3 = make([]int, 4, 20)
	fmt.Println(slice3)
	//for i, j := range slice2 {
	//	fmt.Println(i, j)
	//}
	fmt.Println(slice2)
	ss := append(slice2, 6, 7, 8, 100)
	fmt.Println(ss)
	slice2 = append(slice2, 6, 7, 8, 100, 1919, 2, 2344, 343, 34)
	fmt.Println(slice2)
	fmt.Printf("%p\n", &slice2)
	fmt.Printf("%p\n", &ss)

	fmt.Println(len(slice2[1:]))
	fmt.Println(cap(slice2[1:]))
	slice2[0] = 11
	slice2[1] = 21
	slice2[2] = 31
	slice2[3] = 41
	fmt.Println(slice2)

	//fmt.Print(slice2[1:3])
	//fmt.Println(slice2[:3])

	fmt.Println(len(slice2[1:]))
	fmt.Println(cap(slice2[1:]))

	var s4 []int
	fmt.Println(len(s4))
	fmt.Println(cap(s4))
	s4 = append(s4, 1)
	s4 = append(s4, 2)
	fmt.Println(len(s4))
	fmt.Println(cap(s4))
	s4 = append(s4, 5, 6, 7, 8)
	fmt.Println(len(s4))
	fmt.Println(cap(s4))
	s5 := s4[2:]
	fmt.Println(len(s4))
	fmt.Println(cap(s4))
	fmt.Println(len(s5))
	fmt.Println(cap(s5))

}


*/
