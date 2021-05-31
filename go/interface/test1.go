package main

import "fmt"

func main() {
	i := (interface{})(2)
	fmt.Printf("type:%T, data :%v\n", i, i)
	t := i.(int)
	fmt.Println(t)
}
