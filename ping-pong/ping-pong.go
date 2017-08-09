package main

import (
	"fmt"
	"time"
)

func player(table chan int) {
	for {
		ball := <-table
		fmt.Println(ball)
		ball++
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}

func main() {
	var Ball int
	table := make(chan int)
	for i := 0; i < 100; i++ {
		go player(table)
	}

	table <- Ball
	time.Sleep(1 * time.Second)
	<-table
}
