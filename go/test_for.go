package main

import (
	"fmt"
	"time"
)

func main() {
	for i:=0; i<3; i++ {
		if i == 0 {
			continue
		}
		fmt.Println(i)
		if i == 1 {
			break
		}
	}
	var ch1 chan rune = make(chan rune, 10)
	go func() {
		time.Sleep(10*time.Second)
	}()
	for k,v := range("Go语言") {
		fmt.Printf("%d: %c\n", k, v)
		ch1<-v
	}
	close(ch1)
	time.Sleep(1*time.Second)
	go func() {
		for i:=0; i<6; i++ {
			select {
				case value := <-ch1:
					fmt.Printf("%c", value)
				default:
					fmt.Println("get no thing")
			}
		}
	}()
	time.Sleep(1*time.Second)

	list := []interface{}{'a', "abc", 123, 456}
	for _,value:= range(list) {
		fmt.Println(value)
	}
}