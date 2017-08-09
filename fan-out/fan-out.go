package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(tasksChannel <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		task, ok := <-tasksChannel
		if !ok {
			return
		}
		// All this does is control the printing order
		// WHICH IS BONKERS
		duration := time.Duration(task) * time.Millisecond
		time.Sleep(duration)
		fmt.Println("Processing task", task)
	}
}

func pool(wg *sync.WaitGroup, workers, tasks int) {
	tasksChannel := make(chan int)

	for i := 0; i < workers; i++ {
		go worker(tasksChannel, wg)
	}

	for i := 0; i < tasks; i++ {
		tasksChannel <- i
	}

	close(tasksChannel)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(36)
	go pool(&wg, 36, 60)
	wg.Wait()
}
