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
		time.Sleep(time.Second)
	}
	// 记得 close 信道
	// 不然主函数中遍历完并不会结束，而是会阻塞。
	close(mychan)
}

func main() {
	pipline := make(chan int, 50)

	go fibonacci(pipline)

	for k := range pipline {
		fmt.Println(k)
	}
}
