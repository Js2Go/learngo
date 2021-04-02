package main

import (
	"fmt"
	"sync"
)

func doWorker(id int,
	w worker) {
	for n := range w.in {
		fmt.Printf("Worker %d received %c\n", id, n)
		//go func() { done <- true }()
		w.done()
	}
}

type worker struct {
	in   chan int
	//done chan bool
	//wg *sync.WaitGroup
	done func()
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in:   make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWorker(id, w)
	return w
}

func chanDemo() {
	var workers [10]worker

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	wg.Add(20)

	for i, worker := range workers {
		worker.in <- 'a' + i
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	wg.Wait()
	//
	////wait for all of them
	//for _, worker := range workers {
	//	<-worker.done
	//	<-worker.done
	//}
}

func main() {
	chanDemo()
}
