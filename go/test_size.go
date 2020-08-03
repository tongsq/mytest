package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a byte = 'A'
	var b rune = '我'
	fmt.Printf("byte占用字节数: %d, rune 占用字节数: %d\n", unsafe.Sizeof(a), unsafe.Sizeof(b))
	fmt.Printf("a: %c, b: %c\n", a, b)
	var c string = fmt.Sprintf("%d", a)
	fmt.Println(c, len(c), len("65"))
}
