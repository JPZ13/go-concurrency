package main

import (
	"fmt"
	"time"
)

func producer(channel chan int, duration time.Duration) {
	var i int
	for {
		channel <- i
		i++
		time.Sleep(duration)
	}
}

func reader(output chan int) {
	for x := range output {
		fmt.Println(x)
	}
}

func main() {
	channel := make(chan int)
	output := make(chan int)
	go producer(channel, 100*time.Millisecond)
	go producer(channel, 250*time.Millisecond)
	go reader(output)

	for i := range channel {
		output <- i
	}
}
