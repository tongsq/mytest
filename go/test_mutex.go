package main

import (
	"fmt"
	"sync"
)

func add(count *int, wg *sync.WaitGroup, lock *sync.Mutex) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		lock.Lock()
		*count++
		lock.Unlock()
	}
}

func add2(count *int, wg *sync.WaitGroup, ch chan bool) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		ch <- true
		*count++
		<-ch
	}
}
func main() {
	count := 0
	var wg sync.WaitGroup
	wg.Add(3)
	lock := new(sync.Mutex)
	go add(&count, &wg, lock)
	go add(&count, &wg, lock)
	go add(&count, &wg, lock)

	count2 := 0
	wg.Add(3)
	ch := make(chan bool, 1)
	go add2(&count2, &wg, ch)
	go add2(&count2, &wg, ch)
	go add2(&count2, &wg, ch)
	wg.Wait()
	fmt.Printf("count: %d, count2 : %d", count, count2)
}
