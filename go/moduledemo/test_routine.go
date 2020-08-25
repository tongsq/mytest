package main

import (
	"fmt"
	"moduledemo/util"
	"time"
)

func main() {
	fmt.Println(util.Num)
	go func() {
		for {
			fmt.Println(util.Num)
			time.Sleep(time.Second * 2)
		}
	}()
	go func() {
		for {
			util.Num = util.Num + 1
			time.Sleep(time.Second)
		}
	}()
	time.Sleep(time.Hour)
}
