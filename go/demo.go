package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	i := 0
	for i < 3 {
		fmt.Println(i)
		i++
	}
	defer fmt.Println("defer in main")
	go func() {
		defer fmt.Println(" defer in goroutine")
		panic("test panic")
	}()
	time.Sleep(2 * time.Second)

}
