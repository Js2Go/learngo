package main

import (
	"fmt"
	"time"
)

func main() {
	//var arr [10]int
	for i := 0; i < 10; i++ {
		go func(i int) { // race condition
			for {
				//arr[i]++
				//runtime.Gosched()
				fmt.Printf("Hello from " +
					"goroutine %d\n", i)
			}
		}(i)
	}
	time.Sleep(time.Minute)
	//fmt.Println(arr)
}
