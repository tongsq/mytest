package main

import (
	"fmt"
	"sync"
)

func main() {
	n := 0
	wg := sync.WaitGroup{}
	num := 3000
	wg.Add(num)
	m := sync.Mutex{}
	for i := 0; i < num; i++ {
		go func() {
			m.Lock()
			n++
			m.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(n)
}
