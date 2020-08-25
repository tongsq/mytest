package main

import (
	"fmt"
	"time"
)

func main() {
	index := 1
	go func() {
		for {
			fmt.Println("index is ", index)
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			index = index + 1
			time.Sleep(time.Second)
		}
	}()

	time.Sleep(time.Hour)
}
