package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(name string, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		fmt.Printf("worker : %s, %d\n", name, i)
		time.Sleep(time.Second)
	}
}

func main() {
	//var wg *sync.WaitGroup = new(sync.WaitGroup)
	var wg sync.WaitGroup = sync.WaitGroup{}
	wg.Add(2)
	go worker("协程1", &wg)
	go worker("协程2", &wg)
	wg.Wait()
}
