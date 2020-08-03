package main

import "fmt"

func main() {
	a := new(int)
	fmt.Println(*a)
	mm := make(map[string]int)
	mm["abc"] = 10
	fmt.Println(mm["abc"])
}
