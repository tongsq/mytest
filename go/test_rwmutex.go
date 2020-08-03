package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	lock := new(sync.RWMutex)
	lock.Lock()
	for i := 0; i < 4; i++ {
		go func(i int) {
			fmt.Printf("第 %d 个协程准备开始\n", i)
			lock.RLock()
			defer lock.RUnlock()
			fmt.Printf("第 %d 个协程获取读锁，sleep 1s后释放\n", i)
			time.Sleep(time.Second)
		}(i)
	}
	time.Sleep(time.Second * 2)
	fmt.Println("准备释放写锁...")
	lock.Unlock()
	lock.Lock()
	fmt.Println("exit...")
	lock.Unlock()
}
