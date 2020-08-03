package main

import (
	"fmt"
	"time"
)

func fibonacci(mychan chan int) {
	n := cap(mychan)
	x, y := 1, 1
	for i := 0; i < n; i++ {
		mychan <- x
		x, y = y, x+y
	}
	close(mychan)
	fmt.Println("end, close mychan")
}
func main() {
	pipline := make(chan int, 10)
	go fibonacci(pipline)
	// for k := range pipline {
	// 	fmt.Println(k)
	// 	time.Sleep(1)
	// }
	for {
		v, ok := <-pipline
		fmt.Println(v, ok)
		time.Sleep(time.Second)
		if !ok {
			break
		}
	}
}
