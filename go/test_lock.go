package main

import (
	"fmt"
	"time"
)

func increment(ch chan bool, x *int) {
	ch <- true
	*x++
	<-ch
}

func main() {
	pipline := make(chan bool, 1)
	num := 0
	for i := 0; i < 1000; i++ {
		go increment(pipline, &num)
	}
	time.Sleep(time.Second)
	fmt.Println("num 的值 : ", num)
}
