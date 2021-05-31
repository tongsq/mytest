package main

import (
	"fmt"
	"time"
)

type Demo struct {
	Name string
}

func main() {
	arr := []Demo{
		{
			Name: "aa",
		},
		{
			Name: "bb",
		},
		{
			Name: "cc",
		},
		{
			Name: "dd",
		},
	}
	for _, v := range arr {
		v2 := v
		go func(d *Demo) {
			fmt.Println(d.Name)
		}(&v2)
	}
	time.Sleep(time.Second)
}
