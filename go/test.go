package main

import "fmt"

func main() {
	var i int = 10
	fmt.Println(interface{}(i).(int))
}
