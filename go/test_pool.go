package main

import (
	"fmt"
	"time"
)

var worker_num int

type Pool struct {
	worker chan func()
	size   chan bool
}

func New(size int) *Pool {
	return &Pool{
		worker: make(chan func()),
		size:   make(chan bool, size),
	}
}

func (pool *Pool) workerStart(worker_num int, task func()) {
	defer func() { <-pool.size }()
	fmt.Printf("worder number start:%d \n", worker_num)
	for {
		task()
		task = <-pool.worker
	}
}

func (pool *Pool) RunTask(task func()) {
	select {
	case pool.worker <- task:
	case pool.size <- true:
		go pool.workerStart(worker_num, task)
		worker_num++
	}
}

func main() {
	pool := New(2)
	for i := 0; i < 4; i++ {
		pool.RunTask(func() {
			fmt.Println(time.Now())
			time.Sleep(time.Second * 2)
		})
	}

	time.Sleep(time.Second * 6)
}
