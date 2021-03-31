package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	// 当channel被close掉后，收到的数据会是chan类型的零值
	// 在这为 0
	//for {
	//	n, ok := <-c
	//	if !ok {
	//		break
	//	}
	//	fmt.Printf("Worker %d received %c\n", id, n)
	//}
	for n := range c {
		fmt.Printf("Worker %d received %c\n", id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func chanDemo() {
	//c := make(chan int)
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		//channels[i] = make(chan int)
		channels[i] = createWorker(i)

		channels[i] <- 'a' + i
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)
}

func bufferedChannel() {
	//var channels [10]chan int
	c := make(chan int, 10)
	for i := 0; i < 10; i++ {
		go worker(i, c)

		c <- 'a' + i
		c <- 'A' + i
	}
}

func channelClose() {
	c := make(chan int)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	fmt.Println("Channel as first-class citizen")
	chanDemo()

	fmt.Println("Buffered Channel")
	//bufferedChannel()

	fmt.Println("Channel close")
	//channelClose()
}
