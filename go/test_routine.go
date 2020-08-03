package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(i)
			time.Sleep(2 * time.Second)
			wg.Done()
		}
	}()
	//runtime.Gosched()
	wg.Wait()
}
