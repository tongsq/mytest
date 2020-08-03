package main

import (
	"fmt"
)

func main() {
	i := 1
flag1:
	if i <= 5 {
		fmt.Println(i)
		i++
		goto flag1
	}
	goto flag
	fmt.Println("B")

flag:
	fmt.Println("A")
}
