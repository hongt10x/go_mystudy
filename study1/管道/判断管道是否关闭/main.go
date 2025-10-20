package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()

	for {
		if val, ok := <-ch; ok {
			fmt.Println("Received:", val)
		} else {
			fmt.Println("Channel closed")
			break
		}
	}
}
