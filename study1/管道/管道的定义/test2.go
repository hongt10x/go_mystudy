package main

import "fmt"

func main() {

	quene := make(chan int, 10)
	slice1 := make([]int, 10)
	for i := 0; i < 10; i++ {
		quene <- i
	}
	close(quene)
	for j := 0; j < 10; j++ {
		//fmt.Println(<-quene)
		//slice1 = append(slice1, <-quene)
		slice1[j] = <-quene
	}

	fmt.Println(slice1)
}
