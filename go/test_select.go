package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)
	timeout := make(chan bool, 1)
	c2 <- "hello"
	c1 <- "world"
	go makeTimeout(timeout, 2)
	select {
	case msg := <-c1:
		fmt.Println("c1 receive: ", msg)
	case msg := <-c2:
		fmt.Println("c2 receive: ", msg)
	case <-timeout:
		fmt.Println("timeout")
	default:
		fmt.Println("no data received.")
	}
}

func makeTimeout(ch chan bool, t int) {
	time.Sleep(time.Second * time.Duration(t))
	ch <- true
}
