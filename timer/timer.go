package main

import (
	"fmt"
	"time"
)

func timer(duration time.Duration, display int) <-chan int {
	channel := make(chan int)

	go func() {
		time.Sleep(duration)
		channel <- display
	}()

	return channel
}

func main() {
	for i := 1; i <= 24; i++ {
		outChannel := timer(1*time.Second, i)
		fmt.Println(<-outChannel)
	}
}
