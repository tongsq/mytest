package main

import "fmt"

func main() {
	var num int = 5
	if num:=10; 10 == num {
		num += 10
		fmt.Println(num)
	} else if num > 10 {
		num += 10
		fmt.Println(num)
	} else {
		fmt.Println(num)
	}
	fmt.Println(num)
	names := []string{"Golang", "Java", "Rust", "C"}
	switch name := names[0]; name {
	case "Golang":
	    fmt.Println("A programming language from Google.")
	case "Rust":
	    fmt.Println("A programming language from Mozilla.")
	default:
	    fmt.Println("Unknown!")
	}
	v := 11
	switch i := interface{}(v).(type) {
	case int, int8, int16, int32, int64:
	    fmt.Printf("A signed integer: %d. The type is %T. \n", i, i)
	case uint, uint8, uint16, uint32, uint64:
	    fmt.Printf("A unsigned integer: %d. The type is %T. \n", i, i)
	default:
	    fmt.Println("Unknown!")
	}
}