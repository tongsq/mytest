package main

import (
	"fmt"
	"time"
)
type Sender chan<- string
type Receiver <-chan string
func main() {
	ch1 := make(chan string, 3)
	go func() {
		fmt.Println("start send to ch1")
		ch1 <- "hello"
		fmt.Println("end send to ch1")
	}()
	time.Sleep(time.Duration(2)*time.Second)
	var value string = "receive from ch1: " + <- ch1
	fmt.Println(value)
	close(ch1)

	ch2 := make(chan string, 0)
	var sender Sender = ch2
	var receiver Receiver = ch2
	go func() {
		fmt.Println("start send to ch2")
		sender <- "hello"
		fmt.Println("end send to ch2")
	}()
	time.Sleep(time.Duration(2)*time.Second)
	var value2 string = "receive from ch2: " + <- receiver
	fmt.Println(value2)
	close(ch2)
}