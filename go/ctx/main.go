package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	message := make(chan int, 10)
	for i := 0; i < 10; i++ {
		message <- i
	}

	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()
	defer close(message)
	go func(c context.Context) {
		ticker := time.NewTicker(time.Second)
		for _ = range ticker.C {
			select {
			case <-c.Done():
				fmt.Println("child process exit")
				return
			default:
				fmt.Printf("message num is %d\n", <-message)
			}
		}
	}(c)

	select {
	case <-c.Done():
		fmt.Println("main process exit")
		time.Sleep(time.Second)
	}
}
