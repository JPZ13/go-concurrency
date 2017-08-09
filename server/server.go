package main

import (
	"fmt"
	"net"
	"time"
)

func handler(connection net.Conn, channel chan string) {
	channel <- connection.RemoteAddr().String()
	connection.Write([]byte("ok"))
	connection.Close()
}

func logger(channel chan string) {
	for {
		fmt.Println(<-channel)
	}
}

func server(listener net.Listener, channel chan string) {
	for {
		connection, err := listener.Accept()
		if err != nil {
			continue
		}
		go handler(connection, channel)
	}
}

func main() {
	listener, err := net.Listen("tcp", ":5000")
	if err != nil {
		panic(err)
	}

	channel := make(chan string)
	go logger(channel)
	go server(listener, channel)
	time.Sleep(120 * time.Second)
}
