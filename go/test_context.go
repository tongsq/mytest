package main

import (
	"context"
	"fmt"
	"time"
)

func monitor(ctx context.Context, number int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("监控器%v, 监控结束\n", number)
			return
		default:
			fmt.Printf("监控器%v, 正在监控 %v...\n", number, ctx.Value("test"))
			time.Sleep(time.Second * 2)
		}
	}
}

func main() {
	//ctx01, cancel := context.WithCancel(context.Background())
	ctx01, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*1))
	ctx02, cancel := context.WithDeadline(ctx01, time.Now().Add(time.Second*5))
	ctx02 = context.WithValue(ctx02, "test", "hello")
	defer cancel()

	for i := 1; i <= 5; i++ {
		go monitor(ctx02, i)
	}
	time.Sleep(6 * time.Second)

	if ctx02.Err() != nil {
		fmt.Println("取消的原因:", ctx02.Err())
	}

	fmt.Println("main exit...")
}
