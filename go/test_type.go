package main

import (
	"fmt"
)

func main() {
	var i interface{} = 10
	t1, ok := i.(int)
	fmt.Printf("%d -- %t\n", t1, ok)

	t2, ok := i.(string)
	fmt.Printf("%s -- %t\n", t2, ok)

	var k interface{}
	t3, ok := k.(interface{})
	fmt.Println(t3, ok)

	k = 10
	fmt.Printf("%T\n", k)
	t4, ok := k.(interface{})
	fmt.Printf("%d -- %t\n", t4, ok)

	t5, ok := k.(int)
	fmt.Printf("%d -- %t\n", t5, ok)
	findType(k)
	k = nil
	findType(k)
	var t6 int
	findType(t6)
}

func findType(i interface{}) {
	switch x := i.(type) {
	case int:
		fmt.Println(x, "is int")
	case string:
		fmt.Println(x, "is string")
	case nil:
		fmt.Println(x, "is nil")
	default:
		fmt.Println(x, "not type matched")
	}
}
