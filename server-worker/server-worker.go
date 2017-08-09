package main

import (
	"net"
	"time"
)

func logger(workerChannel, results chan int) {
	for {
		data := <-workerChannel
		data++
		results <- data
	}
}

func parse(results chan int) {
	for {
		<-results
	}
}

func handler(connection net.Conn, channel chan string) {
	address := connection.RemoteAddr().String()
	channel <- address
	time.Sleep(100 * time.Millisecond)
	connection.Write([]byte("ok"))
	connection.Close()
}

func pool(mainChannel chan string, n int) {
	workerChannel := make(chan int)
	results := make(chan int)
	for i := 0; i < n; i++ {
		go logger(workerChannel, results)
	}
	go parse(results)
	for {
		address := <-mainChannel
		length := len(address)
		workerChannel <- length
	}
}

func server(listener net.Listener, mainChannel chan string) {
	for {
		connection, err := listener.Accept()
		if err != nil {
			continue
		}
		go handler(connection, mainChannel)
	}
}

func main() {
	listener, err := net.Listen("tcp", ":5000")
	if err != nil {
		panic(err)
	}
	mainChannel := make(chan string)
	go pool(mainChannel, 4)
	go server(listener, mainChannel)
	time.Sleep(10 * time.Second)
}
