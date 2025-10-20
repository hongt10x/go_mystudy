package main

import (
	"fmt"
	"sync"
)

func main() {
	dict := sync.Map{}
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			dict.Store(i, fmt.Sprintf("test%d...", i))
		}(i)
	}

	wg.Wait()
	// Load
	value, ok := dict.Load(1)
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("cannot find it")
	}
	//	Load and Store

	val, ok := dict.LoadOrStore(200, "GOOD200")
	fmt.Println("val is", val, "ok is", ok)
	val1, ok := dict.LoadOrStore(200, "GOOD200")
	fmt.Println("val is", val1, "ok is", ok)

	//	delete
	dict.Delete(200)
	dict.Delete(1000)
}
